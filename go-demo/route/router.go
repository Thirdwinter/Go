package route

import (
	"github.com/ThirdWinter/Go/go-demo/controllers"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.GET("/user", func(c *gin.Context) {
		c.String(200, "hello,gin")
	})
	user := r.Group("/user")
	{
		user.GET("/info/:id", controllers.UserContraller{}.GetUserInfo)
		user.POST("/list", controllers.UserContraller{}.GetList)
		user.POST("/list/test", controllers.UserContraller{}.GetUserListTest)

		user.POST("/add", controllers.UserContraller{}.AddUesr)
		user.POST("/update", controllers.UserContraller{}.UpdateUser)

		user.POST("/delete", controllers.UserContraller{}.DeleteUser)
	}

	order := r.Group("/order")
	{
		order.POST("/list", controllers.OrderContraller{}.GetList)
	}

	return r
}
