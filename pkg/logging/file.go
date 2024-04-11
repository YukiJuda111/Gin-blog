package logging

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/YukiJuda111/go-gin-blog/pkg/setting"
)

func getLogFilePath() string {
	return setting.AppSetting.LogSavePath
}

func getLogFileFullPath() string {
	prefixPath := getLogFilePath()
	// log20240403.log
	suffixPath := fmt.Sprintf("%s%s.%s", setting.AppSetting.LogSaveName, time.Now().Format(setting.AppSetting.TimeFormat), setting.AppSetting.LogFileExt)
	// runtime/logs/log20240403.log
	return fmt.Sprintf("%s%s", prefixPath, suffixPath)
}

func openLogFile(filePath string) *os.File {
	_, err := os.Stat(filePath)
	switch {
	case os.IsNotExist(err):
		mkDir()
	case os.IsPermission(err):
		log.Fatalln("Permission: ", err)
	}

	// 如果文件不存在，则使用 os.O_CREATE 标志创建它。
	// os.O_APPEND 标志表示如果文件已存在，则在文件末尾追加内容。
	// os.O_WRONLY 标志表示只写模式打开文件。
	handle, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln("Fail to OpenFile: ", err)
	}

	return handle
}

func mkDir() {
	dir, _ := os.Getwd()
	err := os.MkdirAll(dir+"/"+getLogFilePath(), os.ModePerm)
	if err != nil {
		panic(err)
	}
}
