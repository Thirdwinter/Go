package errmsg

const (
	SUCCESS = 200
	ERROR   = 500
	// code =1000... 100开头表示用户模块错误
	ERROR_USERNAME_USED    = 1001 // 用户名以被使用
	ERROR_PASSWORD_WRONG   = 1002 // 用户密码错误
	ERROR_USER_NOT_EXIST   = 1003 // 用户不存在
	ERROR_TOKEN_EXIST      = 1004 // 用户携带token不存在
	ERROR_TOKEN_LONGTIME   = 1005 // 用户携带token超时
	ERROR_TOKEN_WRONG      = 1006 // 用户携带token错误
	ERROR_TOKEN_TYPE_WRONG = 1007 // 用户token格式错误

	// code =2000... 200开头表示文章模块错误
	ERROR_CATENAME_USED = 2001

	// code =3000... 300开头表示分类模块错误
)

var (
	codemsg = map[int]string{
		SUCCESS:                "OK!",
		ERROR:                  "FAIL!",
		ERROR_USERNAME_USED:    "用户名重复!",
		ERROR_PASSWORD_WRONG:   "密码错误!",
		ERROR_USER_NOT_EXIST:   "用户不存在!",
		ERROR_TOKEN_EXIST:      "TOKEN不存在!",
		ERROR_TOKEN_LONGTIME:   "TOKEN超时!",
		ERROR_TOKEN_WRONG:      "TOKEN错误!",
		ERROR_TOKEN_TYPE_WRONG: "TOKEN格式错误",
		ERROR_CATENAME_USED:    "分类已存在",
	}
)

func GetErrMsg(code int) string {
	return codemsg[code]
}
