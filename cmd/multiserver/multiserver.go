package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/gucooing/hkrpg-go/multiserver/db"

	"github.com/gucooing/hkrpg-go/multiserver/config"
	"github.com/gucooing/hkrpg-go/multiserver/multi"
	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/pkg/logger"
)

func main() {
	// 启动读取配置
	confName := "multiserver.json"
	err := config.LoadConfig(confName)
	if err != nil {
		if err == config.FileNotExist {
			p, _ := json.MarshalIndent(config.DefaultConfig, "", "  ")
			cf, _ := os.Create("./conf/" + confName)
			cf.Write(p)
			cf.Close()
			fmt.Printf("找不到配置文件\n已生成默认配置文件 %s \n", confName)
			main()
		} else {
			panic(err)
		}
	}
	appid := alg.GetAppId()
	// 初始化日志
	logger.InitLogger("multiserver"+"["+appid+"]", strings.ToUpper(config.GetConfig().LogLevel))
	logger.Info("hkrpg-go")
	// 初始化
	cfg := config.GetConfig()
	// 初始化数据库
	dbs := db.NewStore(cfg)
	// 初始化服务
	s := multi.NewMulti(cfg, appid, dbs)
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// 启动game
	go func() {
		if err = s.StartMultiServer(); err != nil {
			logger.Error("无法启动game服务器")
		}
	}()

	go func() {
		select {
		case <-done:
			_, cancel := context.WithTimeout(context.Background(), 60*time.Second)
			defer cancel()
			logger.Info("MultiServer 正在关闭")
			s.Close()
			logger.Info("MultiServer 服务已停止")
			logger.CloseLogger()
			os.Exit(0)
		}
	}()
	select {}
}
