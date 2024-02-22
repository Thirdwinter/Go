package controllers

import (
	"strconv"
	"github.com/ThirdWinter/Go/go-demo/models"
	log "github.com/ThirdWinter/Go/mylog"
	"github.com/gin-gonic/gin"
)

type UserContraller struct{}

func init(){
	log.InitLog.Init(true, "Asia/Shanghai", "debug", "./", "log", ".log", 0644, 10*1024)
}

func (u UserContraller) GetUserInfo(c *gin.Context) {
	idstr := c.Param("id")
	name := c.Param("name")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		log.Error("user id return error:%s", err)
		return
	}
	user, _ := models.GetUserTest(id)

	ReturnSucces(c, 0, name, user, 1)
}

func (u UserContraller) AddUesr(c *gin.Context) {
	// var err error
	username := c.DefaultPostForm("username", "")
	idstr := c.DefaultPostForm("id", "9090")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		log.Error("user id return error:%s", err)
		return
	}
	uid, err := models.AddUesr(username, id)
	if err != nil {
		ReturnError(c, 4002, "保存错误")
	}
	ReturnSucces(c, 0, "保存成功", uid, 1)
}

func (u UserContraller) UpdateUser(c *gin.Context) {
	username := c.DefaultPostForm("username", "")
	idstr := c.DefaultPostForm("id", "")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		log.Error("user id return error:%s", err)
		return
	}
	models.UpdateUser(id, username)
	ReturnSucces(c, 0, "更新成功", id, 1)
}

func (u UserContraller) DeleteUser(c *gin.Context) {
	idstr := c.DefaultPostForm("id", "")
	id, _ := strconv.Atoi(idstr)

	err := models.DeleteUser(id)
	if err != nil {
		ReturnError(c, 4002, "删除错误")
		return
	}
	ReturnSucces(c, 0, "删除成功", true, 0)
}

func (u UserContraller) GetUserListTest(c *gin.Context) {
	numstr := c.DefaultPostForm("num", "")
	num, _ := strconv.Atoi(numstr)
	users, err := models.GetUserListTest(num)
	if err != nil || len(users) == 0 {
		ReturnError(c, 404, "无相关数据")
		log.Debug("%#v\n", len(users))
		return
	}
	n := len(users)
	log.Error("%#v\n", users)
	ReturnSucces(c, 200, "获取成功", users, int64(n))
}

func (u UserContraller) GetList(c *gin.Context) {
	//?异常处理,错误返回n3会导致500,加入defer&recover后返回200+空值; 不使前端用户见到err
	// defer func() {
	// 	if err:=recover();err!=nil{
	// 		fmt.Println("捕获异常",err)
	// 	}
	// }()

	// n1:=1
	// n2:=0
	// n3:=n1/n2
	//log.Debug("111")
	ReturnError(c, 404, "user list")
}