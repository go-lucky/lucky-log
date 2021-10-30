// Copyright 2021 Jack lei. All rights reserved.
// Use of this source code is governed by a mit license that can be found
// in the LICENSE file.
package lucky_log

// 错误等级
type Level uint32

// 调试等级最低 ，依次类推，可设置最低显示等级
const (
	TraceLevel Level = iota
	PrintLevel
	InfoLevel
	DebugLevel
	WarningLevel
	ErrorLevel
	PanicLevel
)

// 等级转字符串
var LevelStrMap = map[Level]string{
	TraceLevel:   "trace",
	PrintLevel:   "print",
	InfoLevel:    "info",
	DebugLevel:   "debug",
	WarningLevel: "warning",
	ErrorLevel:   "error",
	PanicLevel:   "panic",
}
