package main

import (
	. "juggle/basic/dao"
	"juggle/basic/middleware"
	_ "juggle/basic/validator"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	dsn := "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	_, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("database connect error: %s", err.Error())
	}

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
