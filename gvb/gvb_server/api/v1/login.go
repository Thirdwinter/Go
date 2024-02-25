package v1

import (
	"github.com/ThirdWinter/Go/gvb_server/middleware"
	"github.com/ThirdWinter/Go/gvb_server/models"
	"github.com/ThirdWinter/Go/gvb_server/utils/errmsg"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context){
	var data models.User
	_ = c.ShouldBindJSON(&data)
	var token string
	var code int
	code=models.CheckLogin(data.Username, data.Password)
	if code == errmsg.SUCCESS{
		token,code=middleware.SetToken(data.Username)
	}
	c.JSON(200, gin.H{
		"status":code,
		"msg":errmsg.GetErrMsg(code),
		"token":token,
	})
}