package cmd

import (
	"encoding/json"
	"github.com/buexplain/lottery/api"
	"github.com/buexplain/lottery/internal/connProcessor"
	"github.com/buexplain/lottery/internal/log"
	"github.com/buexplain/lottery/internal/utils"
	"github.com/buexplain/netsvr-protocol-go/netsvr"
	"google.golang.org/protobuf/proto"
)

// 抽奖
type barrage struct{}

var Barrage = barrage{}

func (r barrage) Init(processor *connProcessor.ConnProcessor) {
	processor.RegisterBusinessCmd(api.Barrage, r.Send)
}

// BarrageParam 弹幕
type BarrageParam struct {
	Message string `json:"message"`
	Id      int    `json:"id"`
}

func (r barrage) Send(_ *netsvr.Transfer, param string, processor *connProcessor.ConnProcessor) {
	payload := BarrageParam{}
	if err := json.Unmarshal(utils.StrToReadOnlyBytes(param), &payload); err != nil {
		log.Logger.Error().Err(err).Msg("Parse BarrageParam failed")
		return
	}
	if payload.Message == "" {
		return
	}
	//广播给所有用户端
	ret := &netsvr.Broadcast{}
	ret.Data = utils.NewResponse(api.Barrage, map[string]interface{}{"code": 0, "message": "弹幕", "data": payload})
	router := &netsvr.Router{}
	router.Cmd = netsvr.Cmd_Broadcast
	router.Data, _ = proto.Marshal(ret)
	pt, _ := proto.Marshal(router)
	processor.Send(pt)
}
