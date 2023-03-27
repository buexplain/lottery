package web

import (
	"github.com/buexplain/lottery/configs"
	"github.com/buexplain/lottery/internal/log"
	"github.com/buexplain/lottery/pkg/wd"
	"html/template"
	"net/http"
	"path"
	"strconv"
	"strings"
)

func getTplData() map[string]any {
	data := map[string]any{}
	//注入参数
	data["conn"] = configs.Config.CustomerWsAddress
	data["handlePattern"] = strings.TrimRight(configs.Config.HandlePattern, "/")
	data["limit3D"] = configs.Config.Limit3D
	data["heartbeatInterval"] = configs.Config.HeartbeatInterval
	data["pageTitle"] = configs.Config.PageTitle
	if configs.Config.WorkerId < 10 {
		data["workerId"] = "00" + strconv.Itoa(int(configs.Config.WorkerId))
	} else if configs.Config.WorkerId < 100 {
		data["workerId"] = "0" + strconv.Itoa(int(configs.Config.WorkerId))
	}
	return data
}

// ClientServer 输出html客户端
func ClientServer() {
	handler := http.FileServer(http.Dir(path.Join(wd.RootPath, "web/asset")))
	handler = http.StripPrefix(configs.Config.HandlePattern, handler)
	http.Handle(configs.Config.HandlePattern, handler)
	tplData := getTplData()
	http.HandleFunc(configs.Config.HandlePattern+"client.html", func(writer http.ResponseWriter, request *http.Request) {
		t, err := template.New("client.html").Delims("{!", "!}").ParseFiles(wd.RootPath + "web/client.html")
		if err != nil {
			log.Logger.Error().Err(err).Msg("模板解析失败")
			return
		}
		err = t.Execute(writer, tplData)
		if err != nil {
			log.Logger.Error().Msgf("模板输出失败：%s", err)
			return
		}
	})
	_ = http.ListenAndServe(configs.Config.ClientListenAddress, nil)
}
