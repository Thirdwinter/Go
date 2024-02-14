package models

import (
	"github.com/ThirdWinter/Go/go-demo/dao"
)

type User struct {
	Id   int
	Username string
}

func (User) TableName() string {
	return "User"
}

func GetUserTest(id int) (User,error) {
	var user User
	err := dao.Db.Where("id=?", id).First(&user).Error
	return user,err
}