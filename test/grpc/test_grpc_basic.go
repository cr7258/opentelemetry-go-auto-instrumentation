// Copyright (c) 2024 Alibaba Group Holding Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"github.com/alibaba/opentelemetry-go-auto-instrumentation/pkg/verifier"
	"go.opentelemetry.io/otel/sdk/trace/tracetest"
	"time"
)

func setupHttp() {
	SetupGRPC()
}

func requestServer() {
	SendReq(context.Background())
}

func main() {
	// starter server
	go setupHttp()
	time.Sleep(3 * time.Second)
	// use a http client to request to the server
	requestServer()
	// verify trace
	verifier.WaitAndAssertTraces(func(stubs []tracetest.SpanStubs) {
		verifier.VerifyRpcClientAttributes(stubs[0][0], "/HelloGrpc/Hello", "grpc", "/HelloGrpc", "Hello")
		verifier.VerifyRpcServerAttributes(stubs[0][1], "/HelloGrpc/Hello", "grpc", "/HelloGrpc", "Hello")
	}, 1)

}