package main

import (
	"fmt"
	"log"
	"syscall"

	"github.com/YukiJuda111/go-gin-blog/models"
	"github.com/YukiJuda111/go-gin-blog/pkg/logging"
	"github.com/YukiJuda111/go-gin-blog/pkg/setting"
	"github.com/YukiJuda111/go-gin-blog/routers"
	"github.com/fvbock/endless"
)

// endless热重启原理，下面这篇博客写的很详细
// https://cloud.tencent.com/developer/article/1848822
func main() {
	setting.Setup()
	models.Setup()
	logging.Setup()

	// router := routers.InitRouter()

	// s := &http.Server{
	// 	Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
	// 	Handler:        router, // http句柄，用于处理程序响应HTTP请求
	// 	ReadTimeout:    setting.ReadTimeout,
	// 	WriteTimeout:   setting.WriteTimeout,
	// 	MaxHeaderBytes: 1 << 20,
	// }

	// s.ListenAndServe()

	endless.DefaultReadTimeOut = setting.ServerSetting.ReadTimeout
	endless.DefaultWriteTimeOut = setting.ServerSetting.WriteTimeout
	endless.DefaultMaxHeaderBytes = 1 << 20
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)

	server := endless.NewServer(endPoint, routers.InitRouter())

	server.BeforeBegin = func(add string) {
		log.Println("Actual pid is ", syscall.Getpid())
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Println("Server err: ", err)
	}
}
