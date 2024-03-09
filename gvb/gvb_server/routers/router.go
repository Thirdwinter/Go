package routers

import (
	v1 "github.com/ThirdWinter/Go/gvb_server/api/v1"
	"github.com/ThirdWinter/Go/gvb_server/global"
	"github.com/ThirdWinter/Go/gvb_server/mdw2"
	"github.com/ThirdWinter/Go/gvb_server/middleware"
	_ "github.com/ThirdWinter/Go/gvb_server/utils/errmsg"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(global.Config.System.Env)
	r := gin.New()
	r.Use(middleware.Logger())
	r.Use(middleware.Cors())
	Auth := r.Group("api/v1")
	//Auth.Use(middleware.JwtToken())
	Auth.Use(mdw2.JwtToken())
	//auth需要使用登录权限认证
	{
		// 用户接口

		Auth.PUT("user/:id", v1.EditUser)      //编辑用户
		Auth.DELETE("user/:id", v1.DeleteUser) //删除用户
		Auth.GET("searchuser", v1.SearchUser)
		Auth.GET("users", v1.GetUsers) //查询用户列表

		// 分类接口
		Auth.POST("category/add", v1.AddCategory)      //添加分类
		Auth.PUT("category/:id", v1.EditCategory)      //修改分类
		Auth.DELETE("category/:id", v1.DeleteCategory) //删除分类
		// 文章接口
		Auth.POST("article/add", v1.AddArticle)      //添加文章
		Auth.PUT("article/:id", v1.EditArticle)      //编辑文章
		Auth.DELETE("article/:id", v1.DeleteArticle) //删除文章
		Auth.POST("upload", v1.UpLoad)               //上传到七牛云
	}
	router := r.Group("api/v1")
	//router不需要登录权限认证
	{
		router.POST("user/add", v1.AddUser)           //注册
		router.POST("login", v1.Login)                //登录
		router.GET("category", v1.GetCategory)        //查询分类列表
		router.GET("article", v1.GetArticle)          //查询文章列表
		router.GET("article/list/:id", v1.GetCateArt) //id查询分类下文章
		router.GET("article/info/:id", v1.GetArtInfo) //查询文章信息

	}

	r.Run(global.Config.System.Addr())

}
