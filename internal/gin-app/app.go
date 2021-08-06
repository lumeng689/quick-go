package main

import (
	"fmt"
	"log"
	"quick-go/pkg/cfg"
	"quick-go/pkg/gredis"
	"quick-go/pkg/logging"
	"syscall"
	"time"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"quick-go/internal/gin-app/router"
)

func init() {
	fmt.Println("global init....")
	// 初始化日志
	logging.SetUp()
	// 初始化配置文件
	cfg.SetUp("./configs/gin-app/app.yaml")
	// 初始化Redis
	gredis.SetUp("localhost:6379")
}

func main() {
	gin.SetMode("debug")
	globalRouter := router.InitRoutes()

	//_ = globalRouter.Run(":8081")

	//server := &http.Server{
	//	Addr:           ":8081",
	//	Handler:        globalRouter,
	//	ReadTimeout:    60 * time.Second,
	//	WriteTimeout:   60 * time.Second,
	//	MaxHeaderBytes: 1 << 20,
	//}

	endless.DefaultReadTimeOut = 60 * time.Second
	endless.DefaultWriteTimeOut = 60 * time.Second
	endless.DefaultMaxHeaderBytes = 1 << 20
	server := endless.NewServer(":8081", globalRouter)

	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
	}

	log.Println("Listen 8081")
	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err: %v", err)
	}
}
