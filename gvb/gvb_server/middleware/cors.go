package middleware

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		cors.New(cors.Config{
			//AllowAllOrigins:true,
			AllowOrigins:  []string{"*"},
			AllowMethods:  []string{"*"},
			AllowHeaders:  []string{"Origin"},
			ExposeHeaders: []string{"Content-Length", "Authorization","Content-Type"},
			//AllowCredentials: true,
			MaxAge: 12 * time.Hour,
		})
	}
}