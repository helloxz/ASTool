package main

import (
	"astool/helper"
	"astool/router"
	"fmt"
	"os"
)

func main() {
	//获取命令行参数
	args := os.Args
	//获取切片长度
	args_len := len(args)

	//如果参数是1，则没有额外参数
	if args_len == 1 {
		fmt.Printf("Please enter parameters!\n")
		os.Exit(1)
	} else if args_len == 2 {
		//启动程序
		if args[1] == "start" {
			// 加载初始化配置文件
			helper.LoadConfig()

			//启动Gin
			router.Start()
		} else if args[1] == "init" {
			//初始化程序，拷贝配置文件，注册服务
			// cli.InitConfig()
		} else if args[1] == "version" {
			// cli.GetVersion()
			// os.Exit(0)
		}
	}
}
