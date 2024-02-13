package controllers

import "github.com/gin-gonic/gin"

type UserContraller struct{}

func (u UserContraller) GetUserInfo(c *gin.Context) {
	id := c.Param("id")
	name := c.Param("name")
	ReturnSucces(c, 0, id, name, 2)
}

func (u UserContraller) GetList(c *gin.Context) {
	ReturnError(c, 404, "error list")
}
