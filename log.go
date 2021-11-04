// Copyright 2021 Jack lei. All rights reserved.
// Use of this source code is governed by a mit license that can be found
// in the LICENSE file.
package lucky_log

import (
	"fmt"
	"path"
	"runtime"
	"strconv"
	"time"

	//"strconv"
	"sync"
	//"time"
)

//record := &MessageRecord{
//Level:         level,
//Message:       message,
//Pid:           os.Getpid(),
//Program:       filepath.Base(os.Args[0]),
//Time:          "",
//FuncName:      runtime.FuncForPC(pc).Name(),
//LongFileName:  file,
//ShortFileName: filepath.Base(file),
//Line:          line,
//Color:         LevelColorFlag[level],
//ColorClear:    LevelColorSeqClear,
//LevelString:   LevelString[level],
//}

// 日志接口
type ILogger interface {
	// 初始化
	InitLogger(config ...interface{}) error
	// 写日志
	/*
		Level:日志级别
		logParams:日志参数
		format: 格式化信息
		args:参数
	*/
	LogMsg(level Level, logogParams map[string]string, format string, args ...interface{}) error

	// 退出
	Exit()
	Flush()
	SetFormatter(format string)
}

type LogMessage struct {
	logLevel      Level
	logFormat     string
	logMsgContent []interface{}
}
type LuckLoggers struct {
	logLuck sync.Mutex
	Loggers []ILogger
}

func (l *LuckLoggers) AppendLogger(logger ILogger) error {
	l.logLuck.Lock()
	defer l.logLuck.Unlock()
	l.Loggers = append(l.Loggers, logger)
	return nil
}
func (l *LuckLoggers) getFormatMsg(f interface{}) string {
	switch f.(type) {
	case string:
		return f.(string)
	default:
		return fmt.Sprint(f)
	}

}
func (l *LuckLoggers) Log(level Level, f interface{}, args ...interface{}) error {
	if len(l.Loggers) == 0 {
		panic("还没有配置日志处理器！")
	} else {
		format := l.getFormatMsg(f)
		//[%(asctime)s] %(filename)s[Line:%(lineno)d] [%(levelname)s]  %(message)s 【filepath:%(pathname)s】
		funcName, file, line, ok := runtime.Caller(2)
		if !ok {
			panic("获取日志信息出错")
		}
		sysLogPars := map[string]string{
			"time":      time.Now().Format("2006-01-02 15:04:05.999"), //time.RFC3339Nano
			"line":      strconv.Itoa(line),
			"levelName": LevelStrMap[level],
			"filePath":  file,
			"fileName":  path.Base(file),
			"funcName":  runtime.FuncForPC(funcName).Name(),
		}

		for _, log := range l.Loggers {
			log.LogMsg(level, sysLogPars, format, args...)
		}
		return nil
	}
}
