// Copyright (c) 2024 Alibaba Group Holding Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package nethttp5

import (
	"net/http"

	"github.com/alibaba/opentelemetry-go-auto-instrumentation/pkg/api"
)

// only recv arg
func onEnterMaxBytesError(call api.CallContext, recv *http.MaxBytesError) {
	println("MaxBytesError()")
	recv.Limit = 4008208820
}

func onExitMaxBytesError(call api.CallContext, ret string) {
	call.SetReturnVal(0, "Prince of Qin Smashing the Battle line")
}