//jwt2 by cookies

package mdw2

import (
	"github.com/ThirdWinter/Go/gvb_server/global"
	"github.com/ThirdWinter/Go/gvb_server/utils/errmsg"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

// 载荷
type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// jwt-key
func Jwt() []byte {
	return []byte(global.Config.System.JwtKey)
}

// 生成token
func SetToken(username string) (string, int) {
	expireTime := time.Now().Add(10 * time.Hour)
	//r_expireTime := time.Now().Add(20 * time.Hour)
	SetClaims := MyClaims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "ginblog",
		},
	}
	// 创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, SetClaims)
	// 生成token
	a_Token, aerr := token.SignedString(Jwt())
	// 生成token错误处理逻辑
	if aerr != nil {
		return "", errmsg.ERROR_TOKEN_CREATE
	}

	return a_Token, errmsg.SUCCESS
}

// 验证atoken
func CheckToken(token string) (*MyClaims, int) {
	claims, err := jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return Jwt(), nil
	})
	if claims == nil {
		return nil, errmsg.ERROR_TOKEN_TYPE_WRONG
	}
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token 已过期
				return nil, errmsg.ERROR_TOKEN_LONGTIME
			}
		}

		// 其他验证错误
		return nil, errmsg.ERROR_TOKEN_WRONG
	}
	return nil, errmsg.SUCCESS
}

// jwt中间件
func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := c.Cookie("token")
		if err != nil {
			c.JSON(401, gin.H{
				"code": 1003,
				"msg":  "Unauthorized",
			})
			c.Abort()
			return
		}
		_, tcode := CheckToken(tokenString)
		switch tcode {
		case 200:
			{
				c.Next()
				return
				//return
			}
		case 4002:
			{
				c.JSON(200, gin.H{
					"code": tcode,
					"msg":  errmsg.GetErrMsg(tcode),
				})
				c.Abort()
				return
			}
		default:
			{
				c.JSON(200, gin.H{
					"code": tcode,
					"msg":  errmsg.GetErrMsg(tcode),
				})
				c.Abort()
				return
			}
		}
	}
}

//"atoken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIiwiZXhwIjoxNzA5NDUxOTk0LCJpc3MiOiJnaW5ibG9nIn0.g6ln3LIv2-gQcNmDQFys0mv5w1Cj5tPEV5j0lkTrGZM",
//"msg": "OK!",
//"rtoken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDk1MjM5ODQsImlzcyI6ImdpbmJsb2cifQ.Kj_LbCkVNXEL086J7WhMRhWToge295iWz_VfiRV1puM",
//"status": 200
