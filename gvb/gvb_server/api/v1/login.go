package v1

import (
	"github.com/ThirdWinter/Go/gvb_server/middleware"
	"github.com/ThirdWinter/Go/gvb_server/models"
	"github.com/ThirdWinter/Go/gvb_server/utils/errmsg"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var data models.User
	_ = c.ShouldBindJSON(&data)
	var atoken string
	var rtoken string
	var code int
	code = models.CheckLogin(data.Username, data.Password)
	if code == errmsg.SUCCESS {
		atoken, rtoken, code = middleware.SetToken(data.Username)
	}
	c.JSON(200, gin.H{
		"code":   code,
		"msg":    errmsg.GetErrMsg(code),
		"atoken": atoken,
		"rtoken": rtoken,
	})
}
