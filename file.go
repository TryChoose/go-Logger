package Logger

import (
	"fmt"
	"os"
	"path"
	"time"
)

type FileLogger struct {
	Level       LogLevel
	filePath    string
	fileName    string
	fileObj     *os.File
	errFile     *os.File
	maxFileSize int64
}

// NewFileLogger 构造函数
func NewFileLogger(level string, filePath string, maxFileSize int64) *FileLogger {
	logLevel, err := ParseLogLevel(level)

	if err != nil {
		panic(err)
	}
	fl := &FileLogger{Level: logLevel,
		filePath:    filePath,
		fileName:    time.Now().Format("2006-01-02") + ".log",
		maxFileSize: maxFileSize,
	}
	err = fl.InitFile()
	if err != nil {
		panic(err)
	}
	return fl
}
func (f *FileLogger) InitFile() error {
	err := os.MkdirAll("logs/right", os.ModePerm)
	if err != nil {
		fmt.Println("Failed to create logs/right directory:", err)
		return err
	}

	err = os.MkdirAll("logs/error", os.ModePerm)
	if err != nil {
		fmt.Println("Failed to create logs/error directory:", err)
		return err
	}
	filepath := path.Join("./logs/right", f.fileName)
	fileObj, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Printf("OPEN LOG FILE FAILED,err :%v", err)
		return err
	}
	//defer fileObj.Close()
	errFileObj, err := os.OpenFile(path.Join("./logs/error", "error.log"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Printf("OPEN ERR LOG FILE FAILED,err :%v", err)
		return err
	}
	//defer errFileObj.Close()
	f.fileObj = fileObj
	f.errFile = errFileObj
	return nil
}

// Enable 志级别开关
func (f *FileLogger) Enable(level LogLevel) bool {
	return level >= f.Level
}
func (f FileLogger) Log(lv LogLevel, format string, a ...any) {
	if f.Enable(lv) {
		msg := fmt.Sprintf(format, a...)
		now := time.Now()
		funcName, fileName, lineno := GetInfo(3)
		if f.CheckSize(f.fileObj) {
			newFile, err := f.SplitFile(f.fileObj)
			if err != nil {
				return
			}
			f.fileObj = newFile
		}
		fmt.Fprintf(f.fileObj, "[%s] [%s][%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:05"), GetLogString(lv), fileName, funcName, lineno, msg)
		if lv >= ERROR {
			if f.CheckSize(f.errFile) {
				newFile, err := f.SplitFile(f.errFile)
				if err != nil {
					return
				}
				f.errFile = newFile
			}
			//如果记录的日志大于等于ERROR级别，还要在err日志文件中记录一下
			fmt.Fprintf(f.errFile, "[%s] [%s][%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:05"), GetLogString(lv), fileName, funcName, lineno, msg)
		}
	}
}

// CheckSize 检查文件大小
func (f *FileLogger) CheckSize(file *os.File) bool {
	fileInfo, err := os.Stat(file.Name())
	if err != nil {
		fmt.Printf("CheckSize :get file Info  failed,err:%v\n", err)
		return false
	}
	return fileInfo.Size() >= f.maxFileSize
}

// SplitFile 切割日志文件
func (f *FileLogger) SplitFile(file *os.File) (*os.File, error) {
	nowStr := time.Now().Format("20060102150405000")
	fileInfo, err := file.Stat()
	if err != nil {
		return nil, fmt.Errorf("SplitFile :get file info failed ,err:%v\n", err)
	}
	//1.关闭当前的日志文件
	file.Close()
	fileName := fileInfo.Name()
	newLogName := fmt.Sprintf("%s-back%s.log", fileName[:len(fileName)-4], nowStr)
	logName := path.Join(f.filePath, fileInfo.Name())
	newLogPath := path.Join(f.filePath, newLogName)

	//2.备份一下 rename
	err = os.Rename(logName, newLogPath)
	if err != nil {
		return nil, fmt.Errorf("SplitFile: failed to rename log file, err: %v\n", err)
	}
	//3.打开一个新的日志文件
	fileObj, err := os.OpenFile(logName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("open file failed err:%v", err)
	}
	//4.将打开的新日志文件对象赋值给 f.fileObj
	return fileObj, nil
}

// Debug 调试
func (f FileLogger) Debug(format string, a ...any) {
	f.Log(DEBUG, format, a...)
}

// Trace 跟踪
func (f FileLogger) Trace(format string, a ...any) {
	f.Log(TRACE, format, a...)
}

// Info 信息
func (f FileLogger) Info(format string, a ...any) {
	f.Log(INFO, format, a...)
}

// Warning 警告
func (f FileLogger) Warning(format string, a ...any) {
	f.Log(WARNING, format, a...)
}

// Error 错误
func (f FileLogger) Error(format string, a ...any) {
	f.Log(ERROR, format, a...)
}

// Fatal 致命错误
func (f FileLogger) Fatal(format string, a ...any) {
	f.Log(FATAL, format, a...)
}

// UnKnow 未知
func (f FileLogger) UnKnow(format string, a ...any) {
	f.Log(UNKNOWN, format, a...)

}

// Close 关闭文件
func (f *FileLogger) Close() {
	f.fileObj.Close()
	f.errFile.Close()
}
