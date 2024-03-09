//package middleware
//
//import (
//	"time"
//
//	"github.com/gin-contrib/cors"
//	"github.com/gin-gonic/gin"
//)
//
//func Cors() gin.HandlerFunc {
//	return cors.New(cors.Config{
//		AllowAllOrigins:  true,
//		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
//		AllowHeaders:     []string{"Origin", "Authorization", "RefreshAuthorization", "Content-Type"},
//		ExposeHeaders:    []string{"Content-Length", "Authorization", "Content-Type"},
//		AllowCredentials: true,
//		MaxAge:           12 * time.Hour,
//	})
//}

// v2 跨域携带cookies
package middleware

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:70"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "RefreshAuthorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Authorization", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})
}
