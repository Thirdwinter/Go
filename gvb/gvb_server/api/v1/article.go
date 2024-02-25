package v1

import (
	"strconv"

	"github.com/ThirdWinter/Go/gvb_server/models"
	"github.com/ThirdWinter/Go/gvb_server/utils/errmsg"
	"github.com/gin-gonic/gin"
)

// AddArticle 添加文章
func AddArticle(c *gin.Context) {
	//检查分类是否存在
	var data models.Article
	_ = c.ShouldBindJSON(&data)
	//code = models.CheckUser(data.Username)
	//if code == errmsg.SUCCESS {
	//	models.CreateUser(&data)
	//}
	//if code == errmsg.ERROR_USERNAME_USED {
	//	code = errmsg.ERROR_USERNAME_USED
	//}
	code = models.CreateArt(&data)
	c.JSON(200, gin.H{
		"status": code,
		"data":   data,
		"msg":    errmsg.GetErrMsg(code),
	})

}

// GetCateArt 查询分类下所有文章
func GetCateArt(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	if pageSize == 0 {
		pageSize = -1 // 为-1时，gorm默认取消Limit分页功能
	}
	if pageNum == 0 {
		pageNum = -1 // 为-1时，gorm默认取消Limit分页功能
	}
	data, code, total := models.GetCateArt(id, pageSize, pageNum)
	c.JSON(200, gin.H{
		"status": code,
		"data":   data,
		"total":  total,
		"msg":    errmsg.GetErrMsg(code),
	})
}

// GetArtInfo 查询文章
func GetArtInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, code := models.GetArtInfo(id)
	c.JSON(200, gin.H{
		"status": code,
		"data":   data,
		"msg":    errmsg.GetErrMsg(code),
	})
}

// GetArticle 查询文章列表
func GetArticle(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	if pageSize == 0 {
		pageSize = -1 // 为-1时，gorm默认取消Limit分页功能
	}
	if pageNum == 0 {
		pageNum = -1 // 为-1时，gorm默认取消Limit分页功能
	}
	data, code, total := models.GetArts(pageSize, pageNum) // 返回一个[]user
	//code = errmsg.SUCCESS
	c.JSON(200, gin.H{
		"status": code,
		"data":   data,
		"total":  total,
		"msg":    errmsg.GetErrMsg(code),
	})
}

// EditArticle 编辑文章
func EditArticle(c *gin.Context) {
	var data models.Article
	id, _ := strconv.Atoi(c.Param("id"))
	err := c.ShouldBindJSON(&data)
	if err != nil {
		return
	}
	code = models.EditArt(id, &data)
	c.JSON(200, gin.H{
		"status": code,
		"msg":    errmsg.GetErrMsg(code),
	})
}

// DeleteArticle 删除文章
func DeleteArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code = models.DeleteArt(id)
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
