package v1

import (
	"strconv"

	"github.com/ThirdWinter/Go/gvb_server/models"
	"github.com/ThirdWinter/Go/gvb_server/utils/errmsg"
	"github.com/gin-gonic/gin"
)

//var code int

// 查询分类是否存在
func CategoryExist(c *gin.Context) {

}

// 添加分类
func AddCategory(c *gin.Context) {
	var data models.Category
	_ = c.ShouldBindJSON(&data)
	code = models.CheckCategory(data.Name)
	if code == errmsg.SUCCESS {
		models.CreateCategory(&data)
	}
	if code == errmsg.ERROR_CATENAME_USED {
		code = errmsg.ERROR_CATENAME_USED
	}
	c.JSON(200, gin.H{
		"status": code,
		//"data":   data,
		"msg":    errmsg.GetErrMsg(code),
	})

}

// 查询分类下所有文章

// 查询分类列表
func GetCategory(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	if pageSize == 0 {
		pageSize = -1 // 为-1时，gorm默认取消Limit分页功能
	}
	if pageNum == 0 {
		pageNum = -1 // 为-1时，gorm默认取消Limit分页功能
	}
	data := models.GetCategory(pageSize, pageNum) // 返回一个[]user
	code = errmsg.SUCCESS
	c.JSON(200, gin.H{
		"status": code,
		"data":   data,
		"msg":    errmsg.GetErrMsg(code),
	})
}

// 编辑分类
func EditCategory(c *gin.Context) {
	var data models.Category
	id,_:=strconv.Atoi(c.Param("id"))
	c.ShouldBindJSON(&data)
	code =models.CheckCategoryEdit(data.Name,int(data.ID))
	if code == errmsg.SUCCESS{
		models.EditCategory(id, &data)
	}
	if code == errmsg.ERROR_CATENAME_USED{
		c.Abort()
	}
	c.JSON(200, gin.H{
		"status":code,
		"msg":errmsg.GetErrMsg(code),
	})
}

// 删除分类
func DeleteCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code = models.DeleteCategory(id)
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
