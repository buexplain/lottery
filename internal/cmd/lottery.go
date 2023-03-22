package cmd

import (
	"github.com/buexplain/lottery/api"
	"github.com/buexplain/lottery/internal/connProcessor"
	"github.com/buexplain/lottery/internal/db"
	"github.com/buexplain/lottery/internal/utils"
	netsvrProtocol "github.com/buexplain/netsvr-protocol-go/protocol"
	"google.golang.org/protobuf/proto"
	"math/rand"
)

// 抽奖
type lottery struct{}

var Lottery = lottery{}

func (r lottery) Init(processor *connProcessor.ConnProcessor) {
	processor.RegisterBusinessCmd(api.LotteryStart, r.Start)
	processor.RegisterBusinessCmd(api.LotteryEnd, r.End)
}

func (r lottery) Start(tf *netsvrProtocol.Transfer, _ string, processor *connProcessor.ConnProcessor) {
	if tf.Session == "" {
		r.forceOffline(tf.UniqId, processor)
		return
	}
	//广播给所有用户端，告知开始抽奖了
	ret := &netsvrProtocol.Broadcast{}
	ret.Data = utils.NewResponse(api.LotteryStart, map[string]interface{}{"code": 0, "message": "开始抽奖", "data": nil})
	router := &netsvrProtocol.Router{}
	router.Cmd = netsvrProtocol.Cmd_Broadcast
	router.Data, _ = proto.Marshal(ret)
	pt, _ := proto.Marshal(router)
	processor.Send(pt)
}

func (r lottery) End(tf *netsvrProtocol.Transfer, _ string, processor *connProcessor.ConnProcessor) {
	if tf.Session == "" {
		r.forceOffline(tf.UniqId, processor)
		return
	}
	//获取所有用户，从中随机一个幸运者
	userList := db.Collect.GetALl()
	i := rand.Intn(len(userList))
	winner := userList[i]
	//广播给所有用户端，告知抽奖成功
	ret := &netsvrProtocol.Broadcast{}
	ret.Data = utils.NewResponse(api.LotteryEnd, map[string]any{"code": 0, "message": "抽奖成功", "data": map[string]any{"winner": winner}})
	router := &netsvrProtocol.Router{}
	router.Cmd = netsvrProtocol.Cmd_Broadcast
	router.Data, _ = proto.Marshal(ret)
	pt, _ := proto.Marshal(router)
	processor.Send(pt)
}

func (lottery) forceOffline(uniqId string, processor *connProcessor.ConnProcessor) {
	ret := &netsvrProtocol.ForceOffline{}
	ret.UniqIds = []string{uniqId}
	router := &netsvrProtocol.Router{}
	router.Cmd = netsvrProtocol.Cmd_ForceOffline
	router.Data, _ = proto.Marshal(ret)
	pt, _ := proto.Marshal(router)
	processor.Send(pt)
}
