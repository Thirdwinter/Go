package controllers

import (
	"strconv"

	"github.com/ThirdWinter/Go/go-demo/models"
	"github.com/ThirdWinter/Go/mylog"
	"github.com/gin-gonic/gin"
)

type UserContraller struct{}

func (u UserContraller) GetUserInfo(c *gin.Context) {
	idstr := c.Param("id")
	name := c.Param("name")
	id,err:=strconv.Atoi(idstr)
	if err!=nil{
		log.Error("user id return error:%s", err)
		return
	}
	user,_:=models.GetUserTest(id)

	ReturnSucces(c, 0,name, user, 1)
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
