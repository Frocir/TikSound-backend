package main

import (
	"TikSound-backend/config"
	"TikSound-backend/server"
	"log"
	"math/rand"
	"os"
	"runtime"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	// 设置核心的最大逻辑数量，以便于最大程度利用宿主机CPU核心
	if len(os.Getenv("GOMAXPROCS")) == 0 {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}
	// 初始化配置
	cfg, err := config.NewConfig()
	if err != nil {
		log.Panicf("Failed to initialize configuration file! err: %#v", err)
		return
	}
	// 初始化日志

	// 初始化服务器
	Server, err := server.NewServer(cfg)
	if err != nil {
		log.Panicf("Failed to initialize service! err: %#v", err)
		return
	}
	// 服务启动
	Server.Run()
}
