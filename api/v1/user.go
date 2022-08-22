/*
 * @Author: TunGuLaMianYG 66915631+TunGuLaMianYG@users.noreply.github.com
 * @Date: 2022-08-18 21:35:44
 * @LastEditors: TunGuLaMianYG 66915631+TunGuLaMianYG@users.noreply.github.com
 * @LastEditTime: 2022-08-22 22:26:17
 * @FilePath: \goblog\api\v1\user.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package v1

import (
	"goblog/module"
	"goblog/utils/errmsg"
	"goblog/utils/snowflake"
	"net/http"

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
	c.ShouldBindJSON(&user)
	code := module.CheckUserByName(user.Username)
	if code == errmsg.SUCCESS {
		userID := snowflake.GenID()
		data := &module.User{
			UserId:   userID,
			Email:    user.Email,
			Username: user.Username,
			Password: user.Password,
			Role:     user.Role,
		}
		code = module.CreateUser(data)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    user,
		"message": errmsg.GetErroMsg(code),
	})
}

// 查询单个用户

// 查询用户列表
func GetUserList(c *gin.Context) {

}

// 编辑用户
func EditUser(c *gin.Context) {
	//
}

// 删除用户
func DeleteUser(c *gin.Context) {
	//
}
