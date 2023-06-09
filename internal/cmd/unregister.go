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

package cmd

import (
	"github.com/buexplain/lottery/internal/connProcessor"
	netsvrProtocol "github.com/buexplain/netsvr-protocol-go/netsvr"
)

type unregister struct{}

var Unregister = unregister{}

func (r unregister) Init(processor *connProcessor.ConnProcessor) {
	processor.RegisterWorkerCmd(netsvrProtocol.Cmd_Unregister, r.UnregisterWorkerOk)
}

func (unregister) UnregisterWorkerOk(_ []byte, processor *connProcessor.ConnProcessor) {
	processor.UnregisterWorkerOk()
}
