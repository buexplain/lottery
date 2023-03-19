package web

import (
	"bytes"
	"compress/gzip"
	"crypto/md5"
	_ "embed"
	"encoding/hex"
	"github.com/buexplain/lottery/configs"
	"github.com/buexplain/lottery/internal/log"
	"github.com/buexplain/lottery/pkg/wd"
	"github.com/rs/zerolog"
	"html/template"
	"net/http"
	"os"
	"strconv"
	"time"
)

//go:embed client.html
var clientSrc string
var clientMd5 string
var client []byte
var clientLastModified string

// 启动的时候就把模板渲染好
func init() {
	if configs.Config.GetLogLevel() == zerolog.DebugLevel {
		return
	}
	t, err := template.New("").Delims("{!", "!}").Parse(clientSrc)
	if err != nil {
		log.Logger.Error().Err(err).Msg("模板解析失败")
		os.Exit(1)
	}
	var buf bytes.Buffer
	var zw *gzip.Writer
	zw, err = gzip.NewWriterLevel(&buf, gzip.DefaultCompression)
	err = t.Execute(zw, getTplData())
	if err != nil {
		log.Logger.Error().Msgf("模板渲染失败：%s", err)
		os.Exit(1)
	}
	if err = zw.Close(); err != nil {
		log.Logger.Error().Msgf("模板压缩失败：%s", err)
		os.Exit(1)
	}
	client = buf.Bytes()
	m := md5.New()
	m.Write(client)
	clientMd5 = hex.EncodeToString(m.Sum(nil))
	clientSrc = ""
	clientLastModified = time.Now().Format(time.RFC1123)
}

func getTplData() map[string]any {
	data := map[string]any{}
	//注入参数
	data["conn"] = configs.Config.CustomerWsAddress
	data["limit3D"] = configs.Config.Limit3D
	data["heartbeatInterval"] = configs.Config.HeartbeatInterval
	if configs.Config.WorkerId < 10 {
		data["workerId"] = "00" + strconv.Itoa(int(configs.Config.WorkerId))
	} else if configs.Config.WorkerId < 100 {
		data["workerId"] = "0" + strconv.Itoa(int(configs.Config.WorkerId))
	}
	return data
}

func clientServerForDev() {
	http.HandleFunc(configs.Config.HandlePattern, func(writer http.ResponseWriter, request *http.Request) {
		t, err := template.New("client.html").Delims("{!", "!}").ParseFiles(wd.RootPath + "web/client.html")
		if err != nil {
			log.Logger.Error().Err(err).Msg("模板解析失败")
			return
		}
		err = t.Execute(writer, getTplData())
		if err != nil {
			log.Logger.Error().Msgf("模板输出失败：%s", err)
			return
		}
	})
	_ = http.ListenAndServe(configs.Config.ClientListenAddress, nil)
}

// ClientServer 输出html客户端
func ClientServer() {
	if configs.Config.GetLogLevel() == zerolog.DebugLevel {
		clientServerForDev()
		return
	}
	http.HandleFunc(configs.Config.HandlePattern, func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Add("Date", time.Now().Format(time.RFC1123))
		writer.Header().Add("Cache-Control", "max-age=604800")
		writer.Header().Add("Last-Modified", clientLastModified)
		writer.Header().Add("ETag", clientMd5)
		if request.Header.Get("If-Modified-Since") == clientLastModified || request.Header.Get("If-None-Match") == clientMd5 {
			http.Error(writer, http.StatusText(http.StatusNotModified), http.StatusNotModified)
			return
		}
		writer.Header().Add("Content-Encoding", "gzip")
		writer.Header().Add("Content-Length", strconv.Itoa(len(client)))
		_, _ = writer.Write(client)
	})
	_ = http.ListenAndServe(configs.Config.ClientListenAddress, nil)
}
