package routers

import (
	v1 "github.com/ThirdWinter/Go/gvb_server/api/v1"
	"github.com/ThirdWinter/Go/gvb_server/global"
	"github.com/ThirdWinter/Go/gvb_server/middleware"
	_ "github.com/ThirdWinter/Go/gvb_server/utils/errmsg"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(global.Config.System.Env)
	r := gin.New()
	r.Use(middleware.Logger())
	Auth := r.Group("api/v1")
	Auth.Use(middleware.JwtToken())
	{
		// 用户接口
		
		Auth.PUT("user/:id", v1.EditUser)
		Auth.DELETE("user/:id", v1.DeleteUser)
		// 分类接口
		Auth.POST("category/add", v1.AddCategory)
		Auth.PUT("category/:id", v1.EditCategory)
		Auth.DELETE("category/:id", v1.DeleteCategory)
		// 文章接口
		Auth.POST("article/add", v1.AddArticle)	
		Auth.PUT("article/:id", v1.EditArticle)
		Auth.DELETE("article/:id", v1.DeleteArticle)
		Auth.POST("upload", v1.UpLoad)
	}
	router:=r.Group("api/v1")
	{
		router.POST("user/add", v1.AddUser)
		router.POST("login", v1.Login)
		router.GET("users", v1.GetUsers)
		router.GET("category", v1.GetCategory)
		router.GET("article", v1.GetArticle)
		router.GET("article/list/:id", v1.GetCateArt)
		router.GET("article/info/:id",v1.GetArtInfo)
		
	}
	

	r.Run(global.Config.System.Addr())

}
