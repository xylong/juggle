package main

import (
	"context"
	. "juggle/basic/src/dao"
	"juggle/basic/src/db"
	"juggle/basic/src/lib"
	"juggle/basic/src/middleware"
	"juggle/basic/src/model"
	_ "juggle/basic/src/validator"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	path, _ := os.Getwd()

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path)

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln(err.Error())
	}

	viper.Unmarshal(lib.Config)

	router := gin.Default()

	// version 1
	v1 := router.Group("v1")
	{
		v1.GET("users", UserList)
		v1.GET("users/:id", db.Cache(UserShow, "id", "user:%s", model.User{}))

		v1.Use(middleware.IsLogin())
		{
			v1.POST("users", UserStore)
			v1.POST("batch-users", UserBatchStore)
			v1.PUT("users/:id", UserUpdate)
			v1.DELETE("users/:id", UserDestroy)
		}
	}

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	// *启动http服务
	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Printf("start server error:%s\n", err.Error())
		}
	}()

	// 初始化数据库
	go func() {
		db.Init()
	}()

	go func() {
		db.InitRedis()
	}()

	lib.ServerNotify()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalln(err.Error())
	}

	log.Println("server exit")
}
