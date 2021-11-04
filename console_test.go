// Copyright 2021 Jack lei. All rights reserved.
// Use of this source code is governed by a mit license that can be found
// in the LICENSE file.
//go test -run console_test.go -v -bench="."
package lucky_log

import (
	"testing"
)

// 性能测试
func BenchmarkLucky(b *testing.B) {
	for i := 0; i < b.N; i++ { //use b.N for looping
		Debug("debug:%s",
			struct {
				UserName     string `json:"user_name"`
				UserPassword string `json:"user_password"`
			}{UserName: "lei", UserPassword: "this is log"})
	}
}

// 并发测试
func BenchmarkCombinationParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Info("info:%s",
				struct {
					UserName     string `json:"user_name"`
					UserPassword string `json:"user_password"`
				}{UserName: "lei", UserPassword: "this is log"})
		}
	})
}

func init() {
	consoleLog := ConsoleLogger{}
	//[2021-10-30 09:42:56,363] etl_queue.py[Line:119] [INFO]
	consoleLog.SetFormatter("[time] [[fileName]:[line]] [[levelName]] [message]")
	AppendLogger(consoleLog)
}
