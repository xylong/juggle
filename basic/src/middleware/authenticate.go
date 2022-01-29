package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// IsLogin 判断是否登录
func IsLogin() gin.HandlerFunc {
	return func(context *gin.Context) {
		if _, ok := context.GetQuery("token"); !ok {
			context.JSON(http.StatusUnauthorized, gin.H{
				"code": 10001,
				"msg":  "未认证",
			})

			context.Abort()
		}

		context.Next()
	}
}
