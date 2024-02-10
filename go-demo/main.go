package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/api/auth/register", func(ctx *gin.Context) {
		// 获取参数
		name:=ctx.PostForm("name")
		phonenumber:=ctx.PostForm("phonenumber")
		password:=ctx.PostForm("password")


		// 数据验证


		// 判断手机号是否存在


		// 创建用户


		// 返回结果
		ctx.JSON(200, gin.H{
			"msg": "注册成功",
		})
	})
	panic(r.Run())
}
