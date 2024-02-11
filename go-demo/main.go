package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/api/auth/register", func(ctx *gin.Context) {
		// 获取参数
		name := ctx.PostForm("name")
		telephone := ctx.PostForm("telephone")
		password := ctx.PostForm("password")

		// 数据验证
		if len(telephone) != 11 {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号必须11位"})
			return
		}

		if len(password) < 6 {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码不能少于6位"})
			return
		}

		if len(name) == 0 {
			name = RandomString(10)
		}
		log.Println(name, telephone, password)
		// 判断手机号是否存在

		// 创建用户

		// 返回结果
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "注册成功",
		})
	})

	panic(r.Run())
}

func RandomString(n int) string {
	var letters = []byte("qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM")
	result := make([]byte, n)
	rng := rand.New(rand.NewSource(time.Now().Unix()))
	for i := range result {
		result[i] = letters[rng.Intn(len(letters))]
	}
	return string(result)
}
