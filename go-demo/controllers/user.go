package controllers

import (
	_ "github.com/ThirdWinter/Go/log"
	"github.com/gin-gonic/gin"
)

type UserContraller struct{}

func (u UserContraller) GetUserInfo(c *gin.Context) {
	id := c.Param("id")
	name := c.Param("name")
	ReturnSucces(c, 0, id, name, 2)
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
