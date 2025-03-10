package hkrpg

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"hkrpg/gdconf"
	"hkrpg/internal"
	"hkrpg/internal/Net"
	"hkrpg/pkg/config"
	"hkrpg/pkg/logger"
)

func Main(configStr string, useDatabase bool, mysqlDsn string) {
	if configStr == "" {
		configByte, _ := json.MarshalIndent(config.DefaultConfig, "", "  ")
		configStr = string(configByte)
	}
	// 启动读取配置
	err := config.LoadConfig(configStr)
	if err != nil {
		//if err == config.FileNotExist {
		//	p, _ := json.MarshalIndent(config.DefaultConfig, "", "  ")
		//	cf, _ := os.Create("/storage/emulated/0/Documents/hkrpg-go/config.json")
		//	cf.Write(p)
		//	cf.Close()
		//	fmt.Printf("找不到配置文件\n已生成默认配置文件 config.json \n")
		//	Main(configStr, useDatabase)
		//} else {
		//	return
		//}
		fmt.Printf("加载配置文件时出错！请尝试重新编译！")
		return
	}
	config.CONF.UseDatabase = useDatabase
	if useDatabase {
		config.CONF.MysqlDsn = mysqlDsn
	} else {
		config.CONF.MysqlDsn = ""
	}
	// 初始化日志
	logger.InitLogger()
	logger.SetLogLevel(strings.ToUpper(config.GetConfig().LogLevel))
	logger.Info("hkrpg-go")

	cfg := config.GetConfig()
	// 初始化
	newserver := internal.NewServer(cfg)
	if newserver == nil {
		logger.Error("服务器初始化失败")
		return
	}
	gdconf.InitGameDataConfig()

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// 启动SDK服务
	go func() {
		if err = newserver.Start(); err != nil {
			logger.Error("无法启动SDK服务器")
		}
	}()

	// 启动game服务
	go func() {
		if err = Net.Run(); err != nil {
			logger.Error("无法启动Game服务器")
		}
	}()

	go func() {
		select {
		case <-done:
			ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
			defer cancel()

			logger.Info("Game服务正在关闭")
			if err = Net.Close(); err != nil {
				logger.Error("无法正常关闭Game服务")
			}
			logger.Info("Game服务已停止")

			logger.Info("SDK服务正在关闭")
			if err = newserver.Shutdown(ctx); err != nil {
				logger.Error("无法正常关闭SDK服务")
			}
			logger.Info("SDK服务已停止")
			logger.CloseLogger()
			os.Exit(0)
		}
	}()
	select {}
}
