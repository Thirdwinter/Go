package models

import (
	"github.com/ThirdWinter/Go/gvb_server/global"
	"github.com/ThirdWinter/Go/gvb_server/utils/errmsg"
	log "github.com/ThirdWinter/Go/mylog"
	"github.com/jinzhu/gorm"
)

type Article struct {
	Category Category `gorm:"foreignkey:Cid"`
	gorm.Model
	Title   string `gorm:"type:varchar(100);not null" json:"title"`
	Cid     int    `gorm:"type:int;not null" json:"cid"`
	Desc    string `gorm:"type:varchar(200)" json:"desc"`
	Content string `gorm:"type:longtext" json:"content"`
	Img     string `gorm:"type:varchar(100)" json:"img"`
}

// 新增文章
func CreateArt(data *Article) (code int) {
	//data.Password =ScryptPw(data.Password)
	err := global.Db.Create(&data).Error
	if err != nil {
		log.Error("create article error: %s", err)
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 查询分类下所有文章
func GetCateArt(id int, pageSize int, pageNum int) ([]Article, int, int) {
	var catArtList []Article
	var total int
	err := global.Db.Preload("Category").Limit(pageSize).Offset((pageNum-1)*pageSize).Where("Cid=?", id).Find(&catArtList).Count(&total).Error
	if err != nil {
		return nil, errmsg.ERROR_CATE_NOT_EXIST, 0
	}
	return catArtList, errmsg.SUCCESS, total
}

// 查询单个文章
func GetArtInfo(id int) (Article, int) {
	var art Article
	err := global.Db.Preload("Category").Where("id=?", id).First(&art).Error
	if err != nil {
		return art, errmsg.ERROR_ART_NOT_EXIST
	}
	return art, errmsg.SUCCESS
}

// 查询文章列表
func GetArts(pageSize int, pageNum int) ([]Article, int, int) {
	var arts []Article
	var total int
	err := global.Db.Preload("Category").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&arts).Count(&total).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errmsg.ERROR, 0
	}
	return arts, errmsg.SUCCESS, total
}

// 编辑文章
func EditArt(id int, data *Article) int {
	var art Article
	// 使用结构体更新，gorm不会更新0值数据
	var maps = make(map[string]interface{})
	maps["title"] = data.Title
	maps["cid"] = data.Cid
	maps["desc"] = data.Desc
	maps["content"] = data.Content
	maps["img"] = data.Img
	err := global.Db.Model(&art).Where("id=?", id).Update(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS

}

// 删除用户
func DeleteArt(id int) int {
	var art Article
	err := global.Db.Where("id=?", id).Delete(&art).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
