// Copyright 2021 Jack lei. All rights reserved.
// Use of this source code is governed by a mit license that can be found
// in the LICENSE file.
// 标准控制台输出
package lucky_log

import (
	"fmt"
	"strings"
)

var consoleFormat string

type ConsoleLogger struct {
}

// 前景 背景 颜色
// ---------------------------------------
// 30  40  黑色
// 31  41  红色
// 32  42  绿色
// 33  43  黄色
// 34  44  蓝色
// 35  45  紫红色
// 36  46  青蓝色
// 37  47  白色
//
// 代码 意义
// -------------------------
//  0  终端默认设置
//  1  高亮显示
//  4  使用下划线
//  5  闪烁
//  7  反白显示
//  8  不可见

// 格式化颜色输出
var levelColorFormatMap = map[Level][2]int{
	TraceLevel:   [2]int{30, 40},
	PrintLevel:   [2]int{30, 40},
	InfoLevel:    [2]int{34, 40},
	DebugLevel:   [2]int{35, 40},
	WarningLevel: [2]int{33, 40},
	ErrorLevel:   [2]int{31, 40},
	PanicLevel:   [2]int{31, 40},


	//TraceLevel:   [2]int{30, 40},
	//PrintLevel:   [2]int{30, 40},
	//InfoLevel:    [2]int{34, 44},
	//DebugLevel:   [2]int{35, 45},
	//WarningLevel: [2]int{33, 43},
	//ErrorLevel:   [2]int{31, 41},
	//PanicLevel:   [2]int{31, 41},

}

func (l ConsoleLogger) InitLogger(config ...interface{}) error {
	return nil
}

func (l ConsoleLogger) SetFormatter(format string) {
	if len(format) > 0 {
		consoleFormat = format
	}
}

func (l ConsoleLogger) LogMsg(level Level, logogParams map[string]string, format string, args ...interface{}) error {
	if _, ok := logogParams["message"]; !ok {
		logogParams["message"] = fmt.Sprintf(format, args...)
	}
	var msgStr string
	msgStr = consoleFormat
	if len(consoleFormat) > 0 {
		for k, v := range logogParams {
			msgStr = strings.ReplaceAll(msgStr, fmt.Sprintf("[%s]", k), v)
		}
	}

	fmt.Printf("%c[1;%d;%dm%s%c[0m\r\n", 0x1B, levelColorFormatMap[level][0], levelColorFormatMap[level][1], msgStr, 0x1B)

	return nil
}

func (l ConsoleLogger) Exit() {

}

func (l ConsoleLogger) Flush() {

}
