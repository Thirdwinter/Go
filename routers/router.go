package routers

import (
	v1 "github.com/ThirdWinter/Go/gvb_server/api/v1"
	"github.com/ThirdWinter/Go/gvb_server/global"
	_ "github.com/ThirdWinter/Go/gvb_server/utils/errmsg"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(global.Config.System.Env)
	r := gin.Default()
	routerv1 := r.Group("api/v1")
	{
		// 用户接口
		routerv1.POST("user/add", v1.AddUser)
		routerv1.GET("users", v1.GetUsers)
		routerv1.PUT("user/:id", v1.EditUser)
		routerv1.DELETE("user/:id", v1.DeleteUser)
		// 分类接口
		routerv1.POST("category/add", v1.AddCategory)
		routerv1.GET("category", v1.GetCategory)
		routerv1.PUT("category/:id", v1.EditCategory)
		routerv1.DELETE("category/:id", v1.DeleteCategory)
	}
	r.Run(global.Config.System.Addr())

}
