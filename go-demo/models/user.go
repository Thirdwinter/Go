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

func AddUesr(usernaem string,id int)(int,error){
	user:=User{Username: usernaem,Id: id}
	err:=dao.Db.Create(&user).Error
	return user.Id,err
}

func UpdateUser(id int, username string){
	dao.Db.Model(&User{}).Where("id=?", id).Update("username",username)
}

func DeleteUser(id int) error{
	err:=dao.Db.Delete(&User{}, id).Error
	return err
}

func GetUserListTest(num int)([]User,error){
	var users []User
	err:=dao.Db.Where("id<?", num).Find(&users).Error
	return users ,err
}