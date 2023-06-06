package Logger

import (
	"fmt"
	"runtime"
	"strings"
)

// Logger 志接口
type Logger interface {
	UnKnow(format string, a ...any)
	Debug(format string, a ...any)
	Trace(format string, a ...any)
	Info(format string, a ...any)
	Warning(format string, a ...any)
	Error(format string, a ...any)
	Fatal(format string, a ...any)
	Enable(lv LogLevel) bool
	Log(lv LogLevel, format string, a ...any)
}

// LogLevel 日志级别
type LogLevel uint16

const (
	UNKNOWN LogLevel = iota
	DEBUG
	TRACE
	INFO
	WARNING
	ERROR
	FATAL
)

// ParseLogLevel 解析日志级别
func ParseLogLevel(s string) (LogLevel, error) {
	s = strings.ToLower(s)
	switch s {
	case "debug":
		return DEBUG, nil
	case "trace":
		return TRACE, nil
	case "info":
		return INFO, nil
	case "waring":
		return WARNING, nil
	case "error":
		return ERROR, nil
	case "fatal":
		return FATAL, nil
	default:
		err := fmt.Errorf("无效的日志级别")
		return UNKNOWN, err
	}
}

// GetLogString 获取日志级别的字符串
func GetLogString(lv LogLevel) string {
	switch lv {
	case DEBUG:
		return "DEBUG"
	case TRACE:
		return "TRACE"
	case INFO:
		return "INFO"
	case WARNING:
		return "WARNING"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	default:
		return "UNKNOWN"
	}
}

// GetInfo 获取函数名，文件名，行号
func GetInfo(skip int) (string, string, int) {
	funcName, file, lineNo, ok := runtime.Caller(skip)
	if !ok {
		fmt.Printf("runtime.Caller() failed\n")
		return "", "", 0
	}
	return runtime.FuncForPC(funcName).Name(), file, lineNo
}
