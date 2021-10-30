// Copyright 2021 Jack lei. All rights reserved.
// Use of this source code is governed by a mit license that can be found
// in the LICENSE file.
package lucky_log

var luckyLog = &LuckLoggers{}

func AppendLogger(logger ILogger) {
	luckyLog.AppendLogger(logger)
}

func Trace(format string, args ...interface{}) {
	luckyLog.Log(TraceLevel, format, args...)
}

func Debug(format string, args ...interface{}) {
	luckyLog.Log(DebugLevel, format, args...)
}

func Print(format string, args ...interface{}) {
	luckyLog.Log(PrintLevel, format, args...)
}

func Info(format string, args ...interface{}) {
	//fmt.Println(format, args)
	luckyLog.Log(InfoLevel, format, args...)


}

func Warning(format string, args ...interface{}) {
	luckyLog.Log(WarningLevel, format, args...)
}

func Error(format string, args ...interface{}) {
	luckyLog.Log(ErrorLevel, format, args...)
}

func Panic(format string, args ...interface{}) {
	luckyLog.Log(PanicLevel, format, args...)
}
