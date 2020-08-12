package main

import (
	"fmt"
	"time"

	"github.com/jiegemena/gotools/timetools"
)

func main() {
	fmt.Println("启动 gorunpy...,5 秒后开始运行")
	for {
		// 5秒检测配置或 是否有服务退出
		time.Sleep(time.Second * 5)
		fmt.Println("读取配置", timetools.DateNowFormatStr())
		if getNewConfig() || getRunStrLen() != runNums {
			fmt.Println("有配置更新", timetools.DateNowFormatStr())
			for k, _ := range allCmd {
				fmt.Println(k)
				if _, ok := getRunStr(allCmd[k].Name); ok {
					fmt.Println(k, allCmd[k].Name, allCmd[k].Path, allCmd[k].Cmd, "exist")
				} else {
					fmt.Println("运行::", k, allCmd[k].Name, allCmd[k].Path, allCmd[k].Cmd)
					setRunStr(allCmd[k].Name, allCmd[k].Name)
					runNums = getRunStrLen()
					go execCommand(allCmd[k].Cmd, allCmd[k].Path, allCmd[k].Name)
				}
			}
		}
	}
}
