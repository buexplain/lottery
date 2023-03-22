/**
* Copyright 2023 buexplain@qq.com
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

package api

type Cmd int32

func (r Cmd) String() string {
	return CmdName[r]
}

// 客户端Cmd路由
const (
	ConnOpen Cmd = iota + 1
	ConnOpenBroadcast
	ConnCloseBroadcast
	LotteryStart
	LotteryEnd
	Barrage
)

var CmdName = map[Cmd]string{
	ConnOpen:           "ConnOpen",           //响应连接成功信息
	ConnOpenBroadcast:  "ConnOpenBroadcast",  //连接成功后，广播给其连接
	ConnCloseBroadcast: "ConnCloseBroadcast", //连接关闭后，广播给其连接
	LotteryStart:       "LotteryStart",       //开始抽奖
	LotteryEnd:         "LotteryEnd",         //结束抽奖
	Barrage:            "Barrage",            //弹幕
}

type Router struct {
	Cmd  Cmd
	Data string
}
