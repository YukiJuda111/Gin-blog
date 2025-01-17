package logging

import (
	"fmt"
	"log"
	"path/filepath"
	"runtime"
)

var (
	DefaultPrefix = ""
	logger        *log.Logger
	levelFlags    = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
)

const (
	DEBUG = iota
	INFO
	WARNING
	ERROR
	FATAL
)

func Setup() {
	filePath := getLogFileFullPath()
	F := openLogFile(filePath)
	log.Println("filePath: ", filePath)
	// 创建一个新的日志记录器。out定义要写入日志数据的IO句柄。
	// prefix定义每个生成的日志行的开头。
	// flag定义了日志记录属性
	logger = log.New(F, DefaultPrefix, log.LstdFlags)
}

func Debug(v ...interface{}) {
	setPrefix(DEBUG)
	logger.Println(v...)
}

func Info(v ...interface{}) {
	setPrefix(INFO)
	logger.Println(v...)
}

func Warn(v ...interface{}) {
	setPrefix(WARNING)
	logger.Println(v...)
}

func Error(v ...interface{}) {
	setPrefix(ERROR)
	logger.Println(v...)
}

func Fatal(v ...interface{}) {
	setPrefix(FATAL)
	logger.Fatalln(v...)
}

func setPrefix(level int) {
	// 使用 runtime.Caller 函数获取调用栈信息。
	// calldepth=2表示调用栈的深度。
	_, file, line, ok := runtime.Caller(2)
	var logPrefix string
	if ok {
		logPrefix = fmt.Sprintf("[%s][%s:%d]", levelFlags[level], filepath.Base(file), line)
	} else {
		logPrefix = fmt.Sprintf("[%s]", levelFlags[level])
	}

	logger.SetPrefix(logPrefix)
}
