/*
 * @Author: TunGuLaMianYG 66915631+TunGuLaMianYG@users.noreply.github.com
 * @Date: 2022-08-16 07:59:32
 * @LastEditors: TunGuLaMianYG 66915631+TunGuLaMianYG@users.noreply.github.com
 * @LastEditTime: 2022-08-22 22:15:14
 * @FilePath: \goblog\main.go
 */
package main

import (
	"context"
	"fmt"
	"goblog/logger"
	"goblog/module"
	"goblog/routes"
	"goblog/utils"
	"goblog/utils/snowflake"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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

	// 初始化 MySql 数据库
	if err := module.MySqlInit(); err != nil {
		fmt.Printf("初始化mysql数据库失败:%v\n", err)
		return
	}
	defer module.MySqlClose()

	// 注册路由
	r := routes.Setup()
	// 启动服务(优雅关机)
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", utils.Conf.Port),
		Handler: r,
	}

	if err := snowflake.Init(utils.Conf.StartTime, utils.Conf.MachineID); err != nil {
		fmt.Printf("init snowflake failed, err:%v\n", err)
		return
	}
	go func() {
		// 开启一个gorroutine启动服务
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen:%s\n", err)
		}
	}()

	// 等待中断信号来优雅地关闭服务器，为关闭服务器操作设置一个5秒的超时
	quit := make(chan os.Signal, 1)
	// kill 默认会发送 syscall.SIGITERM 信号
	// kill -2  发送 syscall.SIGINT 信号，我们常用的 Ctrl+C就是出发系统SIGINT信号
	// kill -9  发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	// signal.Notigy 把收到的 syscall.SIGINT 或 syscall.SIGTERM 信号转发给quit
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
	<-quit                                               // 阻塞在此，当接收到上述两种信号时才会往下执行
	zap.L().Info("关闭服务中")
	// 创建一个5秒超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 5秒内优雅关闭服务(将未处理完的请求处理完再关闭服务)，超过5秒时就超时退出
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Fatal("服务关闭失败：", zap.Error(err))
	}

	zap.L().Info("服务已退出")
}
