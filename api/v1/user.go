/*
 * @Author: TunGuLaMianYG 66915631+TunGuLaMianYG@users.noreply.github.com
 * @Date: 2022-08-18 21:35:44
 * @LastEditors: TunGuLaMianYG 66915631+TunGuLaMianYG@users.noreply.github.com
 * @LastEditTime: 2022-08-24 22:07:18
 * @FilePath: \goblog\api\v1\user.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package v1

import (
	"goblog/module"
	"goblog/utils/errmsg"
	"goblog/utils/snowflake"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var code int

// 查询用户是否存在
func UserExist(c *gin.Context) {
	//
}

// 添加用户
func AddUser(c *gin.Context) {
	var user module.User
	var Data *module.User
	c.ShouldBindJSON(&user)
	code := module.CheckUserByName(user.Username)
	userID := snowflake.GenID()
	if code == errmsg.SUCCESS {
		Data = &module.User{
			UserId:   userID,
			Email:    user.Email,
			Username: user.Username,
			Password: user.Password,
			Role:     user.Role,
		}
		code = module.CreateUser(Data)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    Data,
		"message": errmsg.GetErroMsg(code),
	})
}

// 查询单个用户

// 查询用户列表
func GetUserList(c *gin.Context) {
	pagesize, _ := strconv.Atoi(c.Query("pagesize"))
	pagenum, _ := strconv.Atoi(c.Query("pagenum"))

	if pagesize == 0 {
		pagesize = 10
	}
	if pagenum == 0 {
		pagenum = 1
	}
	users := module.GetUsers(pagesize, pagenum)
	code := errmsg.SUCCESS
	if users == nil {
		code = errmsg.ERROR
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    users,
		"message": errmsg.GetErroMsg(code),
	})
}

// 编辑用户
func EditUser(c *gin.Context) {
	//
}

// 删除用户
func DeleteUser(c *gin.Context) {
	//
}
