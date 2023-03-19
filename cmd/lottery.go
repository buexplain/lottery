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

package main

import (
	"github.com/buexplain/lottery/configs"
	"github.com/buexplain/lottery/internal/cmd"
	"github.com/buexplain/lottery/internal/connProcessor"
	"github.com/buexplain/lottery/internal/log"
	"github.com/buexplain/lottery/pkg/quit"
	"github.com/buexplain/lottery/web"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", configs.Config.WorkerListenAddress)
	if err != nil {
		log.Logger.Error().Msgf("连接服务端失败，%v", err)
		os.Exit(1)
	}
	//启动html客户端的服务器
	go web.ClientServer()
	processor := connProcessor.NewConnProcessor(conn, configs.Config.WorkerId)
	//注册到worker
	if err := processor.RegisterWorker(uint32(configs.Config.ProcessCmdGoroutineNum)); err != nil {
		log.Logger.Debug().Int32("workerId", processor.GetWorkerId()).Err(err).Msg("注册到worker服务器失败")
		os.Exit(1)
	}
	log.Logger.Debug().Int32("workerId", processor.GetWorkerId()).Msg("注册到worker服务器成功")
	//注册各种回调函数
	cmd.ConnSwitch.Init(processor)
	cmd.Lottery.Init(processor)
	//心跳
	go processor.LoopHeartbeat()
	//循环处理worker发来的指令
	for i := 0; i < configs.Config.ProcessCmdGoroutineNum; i++ {
		go processor.LoopCmd()
	}
	//循环写
	go processor.LoopSend()
	//循环读
	go processor.LoopReceive()
	//处理关闭信号
	quit.Wg.Add(1)
	go func() {
		defer func() {
			_ = recover()
			quit.Wg.Done()
		}()
		<-quit.Ctx.Done()
		//取消注册
		processor.UnregisterWorker()
		//优雅关闭
		processor.ForceClose()
	}()
	//开始关闭进程
	select {
	case <-quit.ClosedCh:
		//及时打印关闭进程的日志，避免使用者认为进程无反应，直接强杀进程
		log.Logger.Info().Int("pid", os.Getpid()).Str("reason", quit.GetReason()).Msg("开始关闭进程")
		//通知所有协程开始退出
		quit.Cancel()
		//等待协程退出
		quit.Wg.Wait()
		processor.ForceClose()
		log.Logger.Info().Int("pid", os.Getpid()).Str("reason", quit.GetReason()).Msg("关闭进程成功")
		os.Exit(0)
	}
}
