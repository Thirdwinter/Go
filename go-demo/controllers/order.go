package controllers

import (

	"github.com/gin-gonic/gin"
)

type OrderContraller struct{}

type Search struct {
	Name string `json:"name"`
	Cid  int    `json:"cid"`
}

func (o OrderContraller) GetList(c *gin.Context) {
	//?接收form
	// cid := c.PostForm("cid")
	// name := c.DefaultPostForm("name", "zzx")
	//?接收json
	// param := make(map[string]interface{})
	// err := c.BindJSON(&param)
	//?结构体接收json
	search:=&Search{}
	err:=c.BindJSON(&search)
	
	if err == nil {
		ReturnSucces(c, 0, search.Name,search.Cid, 1)
		return
	}

	ReturnError(c, 404, gin.H{"err": err})
}
