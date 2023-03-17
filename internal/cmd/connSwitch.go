/**
* Copyright 2022 buexplain@qq.com
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
* http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
 */

package cmd

import (
	"github.com/buexplain/lottery/api"
	"github.com/buexplain/lottery/configs"
	"github.com/buexplain/lottery/internal/connProcessor"
	"github.com/buexplain/lottery/internal/db"
	"github.com/buexplain/lottery/internal/log"
	"github.com/buexplain/lottery/internal/utils"
	netsvrProtocol "github.com/buexplain/lottery/pkg/netsvr/protocol"
	"google.golang.org/protobuf/proto"
	"net/url"
	"unicode/utf8"
)

// 连接打开关闭
type connSwitch struct{}

const connTypeAdmin = "admin"

var ConnSwitch = connSwitch{}

func (r connSwitch) Init(processor *connProcessor.ConnProcessor) {
	processor.RegisterWorkerCmd(netsvrProtocol.Cmd_ConnOpen, r.ConnOpen)
	processor.RegisterWorkerCmd(netsvrProtocol.Cmd_ConnClose, r.ConnClose)
}

// ConnOpen 客户端打开连接
func (r connSwitch) ConnOpen(param []byte, processor *connProcessor.ConnProcessor) {
	payload := netsvrProtocol.ConnOpen{}
	if err := proto.Unmarshal(param, &payload); err != nil {
		log.Logger.Error().Err(err).Msg("Parse netsvrProtocol.ConnOpen failed")
		return
	}
	//校验参数
	val, err := url.ParseQuery(payload.RawQuery)
	if err != nil {
		r.forceOffline(payload.UniqId, 1, "参数错误", processor)
		return
	}
	nickname := val.Get("nickname")
	if nickname == "" {
		r.forceOffline(payload.UniqId, 1, "请输入昵称", processor)
		return
	}
	if utf8.RuneCountInString(nickname) > 20 {
		r.forceOffline(payload.UniqId, 1, "昵称太长，最多支持20个字符", processor)
		return
	}
	//如果是管理员登录，则需要判断密码
	connType := val.Get("connType")
	if connType == connTypeAdmin {
		if len(payload.SubProtocol) != 3 {
			r.forceOffline(payload.UniqId, 1, "参数错误", processor)
			return
		}
		//检查加密字符串是否正确，加密值= md5(时间戳+随机值+服务器key)
		if payload.SubProtocol[0] != utils.Md5(payload.SubProtocol[1]+payload.SubProtocol[2]+configs.Config.SecretKey) {
			r.forceOffline(payload.UniqId, 1, "密码错误", processor)
			return
		}
		//检查时间戳是否在合理范围内
		if utils.CheckTimestamp(payload.SubProtocol[1], 10) {
			r.forceOffline(payload.UniqId, 1, "网络超时，请刷新页面再试！", processor)
		}
	}
	//添加到数据库
	user := db.Collect.Add(payload.UniqId, nickname, connType == connTypeAdmin)
	//广播给所有人
	r.broadcast(api.ConnOpenBroadcast, map[string]interface{}{"nickname": nickname, "uniqId": payload.UniqId}, "有新用户登录系统", processor)
	//返回连接成功信息
	data := utils.NewResponse(api.ConnOpen, map[string]any{"code": 0, "message": "登录成功", "data": map[string]any{
		"userList": db.Collect.GetALl(),
		"user":     user,
	}})
	router := &netsvrProtocol.Router{}
	router.Cmd = netsvrProtocol.Cmd_SingleCast
	if connType == connTypeAdmin {
		//管理员要更新网关session
		ret := &netsvrProtocol.InfoUpdate{}
		ret.UniqId = payload.UniqId
		ret.NewSession = nickname
		ret.Data = data
		router.Data, _ = proto.Marshal(ret)
	} else {
		//非管理员，没有session，只是一个游客
		ret := &netsvrProtocol.SingleCast{}
		ret.UniqId = payload.UniqId
		ret.Data = data
		router.Data, _ = proto.Marshal(ret)
	}
	pt, _ := proto.Marshal(router)
	processor.Send(pt)
}

// ConnClose 客户端关闭连接
func (r connSwitch) ConnClose(param []byte, processor *connProcessor.ConnProcessor) {
	payload := netsvrProtocol.ConnClose{}
	if err := proto.Unmarshal(param, &payload); err != nil {
		log.Logger.Error().Err(err).Msg("Parse netsvrProtocol.ConnClose failed")
		return
	}
	user := db.Collect.Del(payload.UniqId)
	if user == nil {
		return
	}
	r.broadcast(api.ConnCloseBroadcast, map[string]interface{}{"nickname": user.Nickname, "uniqId": user.UniqId}, "有用户退出系统", processor)
}

func (connSwitch) forceOffline(uniqId string, code int, message string, processor *connProcessor.ConnProcessor) {
	ret := &netsvrProtocol.ForceOffline{}
	ret.UniqIds = []string{uniqId}
	ret.Data = utils.NewResponse(api.ConnOpen, map[string]interface{}{"code": code, "message": message})
	router := &netsvrProtocol.Router{}
	router.Cmd = netsvrProtocol.Cmd_ForceOffline
	router.Data, _ = proto.Marshal(ret)
	pt, _ := proto.Marshal(router)
	processor.Send(pt)
}

func (connSwitch) broadcast(cmd api.Cmd, data any, message string, processor *connProcessor.ConnProcessor) {
	ret := &netsvrProtocol.Broadcast{}
	ret.Data = utils.NewResponse(cmd, map[string]interface{}{"code": 0, "message": message, "data": data})
	router := &netsvrProtocol.Router{}
	router.Cmd = netsvrProtocol.Cmd_Broadcast
	router.Data, _ = proto.Marshal(ret)
	pt, _ := proto.Marshal(router)
	processor.Send(pt)
}