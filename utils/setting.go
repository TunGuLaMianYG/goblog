/*
 * @Author: TunGuLaMianYG 66915631+TunGuLaMianYG@users.noreply.github.com
 * @Date: 2022-08-17 08:08:05
 * @LastEditors: TunGuLaMianYG 66915631+TunGuLaMianYG@users.noreply.github.com
 * @LastEditTime: 2022-08-17 08:17:13
 * @FilePath: \goblog\utils\setting.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */

package utils

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// Conf 全局变量，用来保存程序的所有配置信息
var Conf = new(AppConfig)

type AppConfig struct {
	Name         string `mapstructure:"name"`
	Mode         string `mapstructure:"mode"`
	Version      string `mapstructure:"version"`
	Port         int    `mapstructure:"port"`
	*LogConfig   `mapstructure:"log"`
	*MySQLConfig `mapstructure:"mysql"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

type MySQLConfig struct {
	Host         string `mapstructure:"host"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	DbName       string `mapstructure:"db_name"`
	Port         int    `mapstructure:"port"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

func Init() (err error) {
	// 指定配置文件名称
	viper.SetConfigFile("config.yaml")
	// 指定配置文件的路径
	viper.AddConfigPath(".")
	// 读取配置文件
	err = viper.ReadInConfig()
	if err != nil {
		// 读取配置信息失败
		fmt.Printf("viper.ReadInConfig() failed,err:%v\n", err)
		return
	}
	if err := viper.Unmarshal(Conf); err != nil {
		// 把读取到的配置信息反序列化到Conf变量中
		fmt.Printf("viper.Unmarshal failed, err:%v\n", err)
	}
	viper.WatchConfig()
	// 监控配置文件的更改信息  fsnotify用来监控文件
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件修改成功")
		if err := viper.Unmarshal(Conf); err != nil {
			fmt.Printf("viper.Unmarshal failed, err:%v\n", err)
		}
	})
	return
}
