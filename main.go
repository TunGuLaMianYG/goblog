/*
 * @Author: TunGuLaMianYG 66915631+TunGuLaMianYG@users.noreply.github.com
 * @Date: 2022-08-16 07:59:32
 * @LastEditors: TunGuLaMianYG 66915631+TunGuLaMianYG@users.noreply.github.com
 * @LastEditTime: 2022-08-16 22:03:15
 * @FilePath: \goblog\main.go
 */
package main

import (
	"fmt"
	"goblog/logger"
	"goblog/utils"

	"go.uber.org/zap"
)

func main() {

	// 加载配置文件
	if err := utils.Init(); err != nil {
		fmt.Printf("初始化配置文件失败：%v\n", err)
		return
	}
	fmt.Println(utils.Conf)
	fmt.Println(utils.Conf.LogConfig == nil)

	// 初始化日志
	if err := logger.Init(utils.Conf.LogConfig); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}
	defer zap.L().Sync()
	zap.L().Debug("Logger init success")

}
