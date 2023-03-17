package web

import (
	"bytes"
	"compress/gzip"
	"crypto/md5"
	_ "embed"
	"encoding/hex"
	"github.com/buexplain/lottery/configs"
	"github.com/buexplain/lottery/internal/log"
	"html/template"
	"os"
	"strconv"
	"time"
)

//go:embed client.html
var ClientMd5 string
var Client []byte
var ClientLastModified string

// 启动的时候就把模板渲染好
func init() {
	t, err := template.New("").Delims("{!", "!}").Parse(ClientMd5)
	if err != nil {
		log.Logger.Error().Err(err).Msg("模板解析失败")
		os.Exit(1)
	}
	data := map[string]interface{}{}
	//注入参数
	data["conn"] = configs.Config.CustomerWsAddress
	data["limit3D"] = configs.Config.Limit3D
	if configs.Config.WorkerId < 10 {
		data["workerId"] = "00" + strconv.Itoa(int(configs.Config.WorkerId))
	} else if configs.Config.WorkerId < 100 {
		data["workerId"] = "0" + strconv.Itoa(int(configs.Config.WorkerId))
	}
	var buf bytes.Buffer
	var zw *gzip.Writer
	zw, err = gzip.NewWriterLevel(&buf, gzip.DefaultCompression)
	err = t.Execute(zw, data)
	if err != nil {
		log.Logger.Error().Msgf("模板渲染失败：%s", err)
		os.Exit(1)
	}
	if err = zw.Close(); err != nil {
		log.Logger.Error().Msgf("模板压缩失败：%s", err)
		os.Exit(1)
	}
	Client = buf.Bytes()
	m := md5.New()
	m.Write(Client)
	ClientMd5 = hex.EncodeToString(m.Sum(nil))
	ClientLastModified = time.Now().Format(time.RFC1123)
}
