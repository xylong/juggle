package main

import (
	. "juggle/basic/dao"
	"juggle/basic/middleware"
	_ "juggle/basic/validator"

	_ "juggle/basic/db"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// version 1
	v1 := router.Group("v1")
	{
		v1.GET("users", UserList)
		v1.GET("users/:id", UserShow)

		v1.Use(middleware.IsLogin())
		{
			v1.POST("users", UserStore)
			v1.POST("batch-users", UserBatchStore)
			v1.DELETE("users/:id", UserDestroy)
		}
	}

	router.Run()
}
