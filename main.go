package main

import (
	"fmt"
	"logtransfer/conf"
	"logtransfer/es"
	"logtransfer/kafka"

	"gopkg.in/ini.v1"
)

// LogTransfer 将日志从kafka中取出发往es

func run() {
	select {}
}

func main() {
	// 0.加载配置文件
	var cfg = new(conf.LogTransferCfg)
	err := ini.MapTo(cfg, "./conf/config.ini")
	if err != nil {
		fmt.Printf("init config failed, err:%v\n", err)
	}
	fmt.Println(cfg)
	// 1.初始化es
	es.Init(cfg.EsCfg.Address, cfg.EsCfg.MaxSize)
	// 2.初始化kafka
	kafka.Init(cfg.KafkaCfg.Address, cfg.KafkaCfg.Topic)
	// 2.1 日志发kafka中取出
	// 2.2 将日志发往es
	run()
}
