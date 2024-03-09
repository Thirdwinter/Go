package v1

import (
	"github.com/ThirdWinter/Go/gvb_server/mdw2"
	"github.com/ThirdWinter/Go/gvb_server/models"
	"github.com/ThirdWinter/Go/gvb_server/utils/errmsg"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var data models.User
	_ = c.ShouldBindJSON(&data)
	var atoken string
	//var rtoken string
	var code int
	code = models.CheckLogin(data.Username, data.Password)
	if code == errmsg.SUCCESS {
		//atoken, rtoken, code = middleware.SetToken(data.Username)
		atoken, _ = mdw2.SetToken(data.Username)
		c.SetCookie("token", atoken, 3600, "/", "", false, true)

		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "Login successful",
		})
		//c.Set("username", data.Username)
		return
	} else {
		c.JSON(200, gin.H{
			"code": 500,
			"msg":  "Login fail",
		})
		c.Abort()
		return
	}
	//c.JSON(200, gin.H{
	//	"code":   code,
	//	"msg":    errmsg.GetErrMsg(code),
	//	"atoken": atoken,
	//	"rtoken": rtoken,
	//})
}
