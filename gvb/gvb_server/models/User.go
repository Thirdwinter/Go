package models

import (
	"encoding/base64"
	"errors"

	"github.com/ThirdWinter/Go/gvb_server/global"
	"github.com/ThirdWinter/Go/gvb_server/utils/errmsg"
	log "github.com/ThirdWinter/Go/mylog"
	_ "github.com/go-playground/validator/v10"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/scrypt"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username" binding:"required" validate:"required,min=4,max=12" label:"用户名"`
	Password string `gorm:"type:varchar(20);not null" json:"password" binding:"required" validate:"required,min=0,max=20" label:"密码"`
	Role     int    `gorm:"type:int;DEFAULT:2" json:"role" validate:"required,gte=2" label:"角色码"`
}

// 用户密码加密
func ScryptPw(password string) string {
	const KeyLen = 10
	salt := []byte{12, 3, 4, 66, 234, 11, 42, 90}
	HashPwd, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, KeyLen)
	if err != nil {
		//!日后处理
		log.Fatal("加盐加密错误：%s", err)
	}
	fpw := base64.StdEncoding.EncodeToString(HashPwd)
	return fpw
}
func (u *User) BeforeSave() {
	u.Password = ScryptPw(u.Password)
} //?gorm 自带

// 注册时查询用户是否存在
func CheckUser(username string) (code int) {
	var users User
	global.Db.Select("id").Where("username=?", username).First(&users)
	if users.ID > 0 {
		return errmsg.ERROR_USERNAME_USED
	}
	return errmsg.SUCCESS
}

// 根据用户名查询用户
func SearchUser(username string) User {
	var user User
	global.Db.Where("username = ?", username).First(&user)
	return user
}

// 更新时检查用户名是否重复
func CheckUserEdit(username string, id int) (code int) {
	var users User
	global.Db.Select("id").Where("username=?", username).First(&users)
	if users.ID > 0 && users.ID != uint(id) {
		return errmsg.ERROR_USERNAME_USED
	}
	return errmsg.SUCCESS
}

// 新增用户
func CreateUser(data *User) (code int) {
	//data.Password =ScryptPw(data.Password)
	err := global.Db.Create(&data).Error
	if err != nil {
		log.Error("create user error: %s", err)
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 查询用户列表
func GetUsers(pageSize int, pageNum int) ([]User, int) {
	var users []User
	var total int
	var err error
	if pageNum <= 0 {
		pageNum = 1
	}
	offset := (pageNum - 1) * pageSize
	err = global.Db.Limit(pageSize).Offset(offset).Find(&users).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return users, 0
		}
		return nil, 0
	}
	err = global.Db.Model(&User{}).Count(&total).Error
	if err != nil {
		return nil, 0
	}
	return users, total
}

// 编辑用户信息 //不能改密码
func EditUser(id int, data *User) int {
	var user User
	// 使用结构体更新，gorm不会更新0值数据
	var maps = make(map[string]interface{})
	maps["username"] = data.Username
	maps["role"] = data.Role
	err := global.Db.Model(&user).Where("id=?", id).Update(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS

}

// 删除用户
func DeleteUser(id int) int {
	var user User
	err := global.Db.Where("id=?", id).Delete(&user).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 登录验证
func CheckLogin(username string, password string) int {
	var user User
	global.Db.Where("username=?", username).First(&user)

	if user.ID == 0 {
		return errmsg.ERROR_USER_NOT_EXIST
	}
	if password == "" {
		return errmsg.ERROR_PASSWORD_NO_EXIST
	}
	encryptedPassword := ScryptPw(password) // 对输入的密码进行加密

	if encryptedPassword != user.Password {
		return errmsg.ERROR_PASSWORD_WRONG
	}
	if user.Role != 1 {
		return errmsg.ERROR_USER_NO_RIGHT
	}
	return errmsg.SUCCESS
}
