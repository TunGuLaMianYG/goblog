/*
 * @Author: TunGuLaMianYG 66915631+TunGuLaMianYG@users.noreply.github.com
 * @Date: 2022-08-17 21:39:52
 * @LastEditors: TunGuLaMianYG 66915631+TunGuLaMianYG@users.noreply.github.com
 * @LastEditTime: 2022-08-18 08:15:39
 * @FilePath: \goblog\model\mysql.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package module

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var db *gorm.DB
var err error

// MySql 初始化
func MySqlInit() (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		viper.GetString("mysql.user"),
		viper.GetString("msql.password"),
		viper.GetString("mysql.host"),
		viper.GetInt("mysql.port"),
		viper.GetString("mysql.dbname"),
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// 禁用默认表名的复数形式
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		zap.L().Error("连接数据库失败:%v\n", zap.Error(err))
		return
	}

	// 自动迁移
	db.AutoMigrate(&User{}, &Article{}, &Category{})

	sqldb, _ := db.DB()

	// 设置连接池中的最大限制连接数
	sqldb.SetMaxIdleConns(viper.GetInt("max_idle_conns"))
	// 设置数据库的最大连接数量
	sqldb.SetMaxOpenConns(viper.GetInt("max_open_conns"))
	// 设置连接的最大可复用时间
	sqldb.SetConnMaxLifetime(10 * time.Second)
	return
}

// 关闭db连接
func MySqlClose() {
	sqlDB, _ := db.DB()
	_ = sqlDB.Close()
}
