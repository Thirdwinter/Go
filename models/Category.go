package models

import (
	"github.com/ThirdWinter/Go/gvb_server/global"
	"github.com/ThirdWinter/Go/gvb_server/utils/errmsg"
	log "github.com/ThirdWinter/Go/mylog"
	"github.com/jinzhu/gorm"
)

type Category struct {
	ID   uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}

// 注册时查询分类是否存在
func CheckCategory(Name string) (code int) {
	var cate Category
	global.Db.Select("id").Where("Name=?", Name).First(&cate)
	if cate.ID > 0 {
		return errmsg.ERROR_CATENAME_USED
	}
	return errmsg.SUCCESS
}

// 更新时检查分类是否重复
func CheckCategoryEdit(Name string, ID int) (code int) {
	var cate Category
	global.Db.Select("ID").Where("Name=?", Name).First(&cate)
	if cate.ID > 0 && cate.ID != uint(ID) {
		return errmsg.ERROR_CATENAME_USED
	}
	return errmsg.SUCCESS
}

// 新增分类
func CreateCategory(data *Category) (code int) {
	err := global.Db.Create(&data).Error
	if err != nil {
		log.Error("create Category error: %s", err)
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 查询分类列表
func GetCategory(pageSize int, pageNum int) []Category {
	var cates []Category
	err := global.Db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&cates).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return cates
}

// 编辑分类信息
func EditCategory(ID int, data *Category) int {
	var cate Category
	// 使用结构体更新，gorm不会更新0值数据
	var maps = make(map[string]interface{})
	maps["name"] = data.Name
	err := global.Db.Model(&cate).Where("ID=?", ID).Update(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS

}

// 删除分类
func DeleteCategory(ID int) int {
	var cate Category
	err := global.Db.Where("ID=?", ID).Delete(&cate).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//!查询分类下所有文章
