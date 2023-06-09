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

package configs

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/buexplain/lottery/pkg/wd"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"path/filepath"
	"strings"
)

type config struct {
	//日志级别 debug、info、warn、error
	LogLevel string
	//worker服务的监听地址
	WorkerListenAddress string
	//网关服务唯一编号，如果配置不对，网关会拒绝连接
	ServerId uint32
	//输出客户端界面的http服务的监听地址
	ClientListenAddress string
	//url路由
	HandlePattern string
	//客户端心跳间隔，单位秒
	HeartbeatInterval uint8
	//customer服务的websocket连接地址
	CustomerWsAddress string
	//让worker为我开启n条协程来处理我的请求
	ProcessCmdGoroutineNum int
	//业务进程注册到网关的工作id
	WorkerId int32
	//加密的key
	SecretKey string
	//前端3d球的渲染数限制
	Limit3D uint16
	//页面标题
	PageTitle string
	//欢迎语言
	PageWelcome string
}

func (r *config) GetLogLevel() zerolog.Level {
	switch r.LogLevel {
	case "debug":
		return zerolog.DebugLevel
	case "info":
		return zerolog.InfoLevel
	case "warn":
		return zerolog.WarnLevel
	case "error":
		return zerolog.ErrorLevel
	}
	return zerolog.ErrorLevel
}

var Config *config

func init() {
	var configFile string
	flag.StringVar(&configFile, "config", filepath.Join(wd.RootPath, "configs/lottery.toml"), "Set lottery.toml file")
	flag.Parse()
	//读取配置文件
	c, err := os.ReadFile(configFile)
	if err != nil {
		log.Printf("Read lottery.toml failed：%s", err)
		os.Exit(1)
	}
	//解析配置文件到对象
	Config = new(config)
	if _, err = toml.Decode(string(c), Config); err != nil {
		log.Printf("Parse lottery.toml failed：%s", err)
		os.Exit(1)
	}
	if Config.ProcessCmdGoroutineNum <= 0 {
		Config.ProcessCmdGoroutineNum = 1
	}
	Config.HandlePattern = strings.TrimRight(Config.HandlePattern, "/") + "/"
}
