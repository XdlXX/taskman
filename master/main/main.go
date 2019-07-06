package main

import (
	"flag"
	"fmt"
	"runtime"
	"start.jyz/taskman/master"
)

var (
	confFile string
)

// 接受命令行参数
func initArgs() {
	// master -config ./master.json
	flag.StringVar(&confFile, "config", "./master.json", "传入配置文件路径")
	flag.Parse()
}

// 初始化线程数量
func initEnv() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	var (
		err error
	)
	// 初始化命令行参数
	initArgs()
	// 初始化线程
	initEnv()
	// 加载配置
	if err = master.InitConfig(confFile); err != nil {
		goto ERR
	}
	// 启动API HTTP 服务
	if err = master.InitApiServer(); err != nil {
		goto ERR
	}

	// 正常退出
	return
	// 异常退出
ERR:
	fmt.Println(err)
}
