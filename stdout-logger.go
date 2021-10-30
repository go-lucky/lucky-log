// Copyright 2021 Jack lei. All rights reserved.
// Use of this source code is governed by a mit license that can be found
// in the LICENSE file.
// 标准控制台输出
package lucky_log

import (
	"fmt"
	"strings"
)

var  stdoutFormat string
type StdoutLogger struct {

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

func (l StdoutLogger) InitLogger(config ...interface{}) error {
	return nil
}

func (l StdoutLogger) SetFormatter(format string) {
	if len(format) > 0 {
		stdoutFormat = format
	}
}

// todo 未完成
func (l StdoutLogger) LogMsg(level Level, logogParams map[string]string, format string, args ...interface{}) error {
	if _, ok := logogParams["message"]; !ok {
		logogParams["message"] = fmt.Sprintf(format, args...)
	}
	var msgStr string
	msgStr = stdoutFormat
	if len(stdoutFormat) > 0 {
		for k, v := range logogParams {
			msgStr = strings.ReplaceAll(msgStr, fmt.Sprintf("[%s]", k), v)
		}
	}
	//fmt.Printf("====1",msgStr,"====\n\n")
	//msgStr = fmt.Sprintf(levelColorFormatMap[level], msgStr)
	//fmt.Printf("====2",msgStr,"====\n\n")
	fmt.Printf("%c[1;%d;%dm%s%c[0m\r\n", 0x1B, levelColorFormatMap[level][0], levelColorFormatMap[level][1], msgStr, 0x1B)
	//fmt.Println()
	//fmt.Printf("%c[%d;%d;%dm%s%c[0m", 0x1B, 1, 31, 40, msgStr, 0x1B)
	//fmt.Print("0x1B[1;30;40m世界你好0x1B[0m")

	//os.Stdout.WriteString(msgStr)
	return nil
}

func (l StdoutLogger) Exit() {

}

func (l StdoutLogger) Flush() {

}
