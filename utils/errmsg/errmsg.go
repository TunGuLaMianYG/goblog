package errmsg

const (
	SUCCESS = 200
	ERROR   = 500

	// code = 1000... 用户模块的错误
	ERROR_USERNAME_USED    = 1001 // 用户名被使用了
	ERROR_PASSWORD_WRONG   = 1002 // 用户密码错误
	ERROR_USER_NOT_EXIS    = 1003 // 用户不存在
	EROOR_TOKEN_EXIST      = 1004 // token不存在
	ERROR_TOKEN_RUNTIME    = 1005 // token超时
	ERROR_TOKEN_WRONG      = 1006 // 错误的token
	ERROR_TOKEN_TYPE_WRONG = 1007 // token格式错误

	// code = 2000... 文章模块的错误

	// code = 3000... 分类模块的错误
)

var CodeMsg = map[int]string{
	SUCCESS:                "OK",
	ERROR:                  "FAIL",
	ERROR_USERNAME_USED:    "用户名已存在！",
	ERROR_PASSWORD_WRONG:   "密码错误",
	ERROR_USER_NOT_EXIS:    "用户不存在",
	EROOR_TOKEN_EXIST:      "TOKEN不存在",
	ERROR_TOKEN_RUNTIME:    "TOKEN已过期",
	ERROR_TOKEN_WRONG:      "TOKEN不正确",
	ERROR_TOKEN_TYPE_WRONG: "TOKEN格式错误",
}

func GetErroMsg(code int) string {

	return CodeMsg[code]
}
