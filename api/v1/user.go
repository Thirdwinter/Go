package v1

import (
	"strconv"

	"github.com/ThirdWinter/Go/gvb_server/models"
	"github.com/ThirdWinter/Go/gvb_server/utils/errmsg"
	"github.com/gin-gonic/gin"
)

var code int

// 查询用户是否存在
func UserExist(c *gin.Context) {

}

// 添加用户
func AddUser(c *gin.Context) {
	var data models.User
	_ = c.ShouldBindJSON(&data)
	code = models.CheckUser(data.Username)
	if code == errmsg.SUCCESS {
		models.CreateUser(&data)
	}
	if code == errmsg.ERROR_USERNAME_USED {
		code = errmsg.ERROR_USERNAME_USED
	}
	c.JSON(200, gin.H{
		"status": code,
		"data":   data,
		"msg":    errmsg.GetErrMsg(code),
	})

}

// 查询用户

// 查询用户列表
func GetUsers(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	if pageSize == 0 {
		pageSize = -1 // 为-1时，gorm默认取消Limit分页功能
	}
	if pageNum == 0 {
		pageNum = -1 // 为-1时，gorm默认取消Limit分页功能
	}
	data := models.GetUsers(pageSize, pageNum) // 返回一个[]user
	code = errmsg.SUCCESS
	c.JSON(200, gin.H{
		"status": code,
		"data":   data,
		"msg":    errmsg.GetErrMsg(code),
	})
}

// 编辑用户
func EditUser(c *gin.Context) {
	var data models.User
	id,_:=strconv.Atoi(c.Param("id"))
	c.ShouldBindJSON(&data)
	code =models.CheckUserEdit(data.Username,id)
	if code == errmsg.SUCCESS{
		models.EditUser(id, &data)
	}
	if code == errmsg.ERROR_USERNAME_USED{
		c.Abort()
	}
	c.JSON(200, gin.H{
		"status":code,
		"msg":errmsg.GetErrMsg(code),
	})
}

// 删除用户
func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code = models.DeleteUser(id)
	if id == 0 {
		c.JSON(200, gin.H{
			"status": errmsg.ERROR,
			"msg":    "请求错误！",
		})
		return
	}

	c.JSON(200, gin.H{
		"status": code,
		"data":   "",
		"msg":    errmsg.GetErrMsg(code),
	})
}
