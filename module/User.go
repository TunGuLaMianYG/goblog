/*
 * @Author: TunGuLaMianYG 66915631+TunGuLaMianYG@users.noreply.github.com
 * @Date: 2022-08-17 21:38:40
 * @LastEditors: TunGuLaMianYG 66915631+TunGuLaMianYG@users.noreply.github.com
 * @LastEditTime: 2022-08-30 08:15:03
 * @FilePath: \goblog\model\User.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package module

import (
	"encoding/base64"
	"fmt"
	"goblog/utils/errmsg"
	"time"

	"go.uber.org/zap"
	"golang.org/x/crypto/scrypt"
	"gorm.io/gorm"
)

var code int

type User struct {
	// 导入gorm默认结构
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	UserId    int64  `gorm:"column:user_id;type:bigint(20);NOT NULL" json:"user_id" binding:"required"`
	Username  string `gorm:"column:username;type:varchar(64);NOT NULL" json:"username" binding:"required"`
	Password  string `gorm:"column:password;type:varchar(64);NOT NULL" json:"password"`
	Email     string `gorm:"column:email;type:varchar(64)" json:"email"`
	Role      int    `gorm:"type:int" json:"role"`
	// Username string `gorm:"type:varchar(20);not null" json:"username"`
	// Password string `gorm:"type:varchar(20);not null" json:"password"`
	// // 角色设置
	// Role int `gorm:"type:int" json:"role"`
	//Id       int64  `gorm:"column:id;type:bigint(20);primary_key;AUTO_INCREMENT" json:"id"`
	// Gender     int          `gorm:"column:gender;type:tinyint(4);default:0;NOT NULL" json:"gender"`
	// CreateTime time.Time `gorm:"column:create_time;type:timestamp;default:CURRENT_TIMESTAMP" json:"create_time"`
	// UpdateTime time.Time `gorm:"column:update_time;type:timestamp;default:CURRENT_TIMESTAMP" json:"update_time"`

}

// 查询用户是饭否存在
func CheckUserByName(name string) int {
	var users User
	//fmt.Println(name)
	fmt.Println(&users.Username)
	db.Select("id").Where("username = ?", name).First(&users)
	fmt.Println(users)
	if users.ID > 0 {
		return errmsg.ERROR_USERNAME_USED
	}
	return errmsg.SUCCESS
}

// 新增用户
func CreateUser(data *User) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 查询用户列表
func GetUsers(pagesize, pagenum int) []User {
	users := []User{}
	err = db.Limit(pagesize).Offset((pagenum - 1) * pagesize).Find(&users).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return users
}

// 编辑用户

// 删除用户

// 密码加密
func ScryptPw(password string) string {
	const KeyLen = 10
	salt := make([]byte, 8)
	salt = []byte{12, 32, 4, 6, 66, 22, 222, 11}

	HashPw, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, KeyLen)
	if err != nil {
		zap.L().Error("密码加密失败：%v\n", zap.Error(err))
	}
	fpw := base64.StdEncoding.EncodeToString(HashPw)
	return fpw
}
