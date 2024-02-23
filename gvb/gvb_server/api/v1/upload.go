package v1

import (
	"github.com/ThirdWinter/Go/gvb_server/service"
	"github.com/ThirdWinter/Go/gvb_server/utils/errmsg"
	"github.com/gin-gonic/gin"
)

func UpLoad(c *gin.Context) {
	file,fileHeader,_:=c.Request.FormFile("file")
	filesize:=fileHeader.Size
	url,code:=service.UpLoadFile(file,filesize)
	c.JSON(200,gin.H{
		"status":code,
		"msg":errmsg.GetErrMsg(code),
		"url":url,
	})
}