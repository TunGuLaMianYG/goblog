package v1

import (
	"fmt"
	"goblog/module"
	"goblog/utils/errmsg"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 查询用户是否存在
func UserExist(c *gin.Context) {
	//
}

// 添加用户
func AddUser(c *gin.Context) {
	var user module.User
	err := c.ShouldBindJSON(&user)
	fmt.Println(err)
	if err == nil {
		fmt.Println(user.Username)
		code := module.CheckUserByName(user.Username)
		if code == errmsg.SUCCESS {
			code = module.CreateUser(&user)
		}
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"data":    user,
			"message": errmsg.GetErroMsg(code),
		})
	}

}

func Live(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
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
