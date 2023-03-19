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

package utils

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"github.com/buexplain/lottery/api"
	"io"
	"strconv"
	"time"
	"unsafe"
)

// NewResponse 构造一个返回给客户端响应
func NewResponse(cmd api.Cmd, data interface{}) []byte {
	tmp := map[string]interface{}{"cmd": cmd, "data": data, "version": time.Now().UnixMilli()}
	ret, _ := json.Marshal(tmp)
	return ret
}

func Md5(s string) string {
	h := md5.New()
	_, err := io.WriteString(h, s)
	if err != nil {
		return ""
	}
	return hex.EncodeToString(h.Sum(nil))
}

func CheckTimestamp(t string, scope int64) bool {
	i, err := strconv.ParseInt(t, 10, 0)
	if err != nil {
		return false
	}
	return time.Now().Unix()-i < scope
}

// StrToReadOnlyBytes 字符串无损转字节切片，转换后的切片，不能做修改操作，因为go的字符串是不可修改的
func StrToReadOnlyBytes(s string) []byte {
	if s == "" {
		return nil
	}
	return unsafe.Slice(unsafe.StringData(s), len(s))
}
