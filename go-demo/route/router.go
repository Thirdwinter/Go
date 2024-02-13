package route

import "github.com/gin-gonic/gin"

func Router() *gin.Engine{
	r := gin.Default()
	r.GET("/user",func(c *gin.Context) {
		c.String(200,"hello,gin")
	})
	user:=r.Group("/user")
	{
		user.POST("/list",func(ctx *gin.Context) {
			ctx.String(200,"list")
			ctx.JSON(200,"userlist")
		})

		user.PUT("/add",func(ctx *gin.Context) {
			ctx.String(200,"put")
		})

		user.DELETE("/delete",func(ctx *gin.Context) {
			ctx.String(200,"delete")
		})
	}
	return r
}