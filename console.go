package Logger

import (
	"fmt"
	"time"
)

// ConsoleLogger 控制台日志结构体
type ConsoleLogger struct {
	Lever LogLevel
}

func NewConsoleLogger(leverStr string) *ConsoleLogger {
	level, err := ParseLogLevel(leverStr)
	if err != nil {
		panic(err)
	}
	return &ConsoleLogger{Lever: level}
}

func (c ConsoleLogger) Enable(lv LogLevel) bool {
	return lv >= c.Lever
}

// Debug 调试
func (c ConsoleLogger) Debug(format string, a ...any) {
	c.Log(DEBUG, format, a...)

}

// Trace 跟踪
func (c ConsoleLogger) Trace(format string, a ...any) {
	c.Log(TRACE, format, a...)
}

// Info 信息
func (c ConsoleLogger) Info(format string, a ...any) {
	c.Log(INFO, format, a...)

}

// Warning 警告
func (c ConsoleLogger) Warning(format string, a ...any) {
	c.Log(WARNING, format, a...)

}

// Error  错误
func (c ConsoleLogger) Error(format string, a ...any) {

	c.Log(ERROR, format, a...)
}

// Fatal 致命错误
func (c ConsoleLogger) Fatal(format string, a ...any) {
	c.Log(FATAL, format, a...)

}

// UnKnow 未知
func (c ConsoleLogger) UnKnow(format string, a ...any) {
	c.Log(UNKNOWN, format, a...)

}

// Log 日志
func (c ConsoleLogger) Log(lv LogLevel, format string, a ...any) {
	if c.Enable(lv) {
		msg := fmt.Sprintf(format, a...)
		now := time.Now()
		funcName, fileName, lineno := GetInfo(3)
		fmt.Printf("[%s] [%s][%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:05"), GetLogString(lv), fileName, funcName, lineno, msg)
	}
}
