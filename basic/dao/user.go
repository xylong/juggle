package dao

import (
	"github.com/gin-gonic/gin"
	"juggle/basic/model"
	"net/http"
	"strconv"
)

func UserList(ctx *gin.Context)  {
	query:=&model.UserQuery{}
	if err:=ctx.BindQuery(query);err!=nil {
		ctx.JSON(http.StatusBadRequest,gin.H{
			"code": 10000,
			"msg":"error:"+err.Error(),
		})

		return
	}

	

	ctx.JSON(http.StatusOK,gin.H{
		"code":0,
		"msg":"",
		"data":query,
	})
}

func UserShow(ctx *gin.Context)  {
	id,_:=strconv.Atoi(ctx.Param("id"))
	ctx.JSON(http.StatusOK,gin.H{
		"code":0,
		"msg":"",
		"data":model.NewUser(id,"张三"),
	})
}

func UserStore(ctx *gin.Context)  {
	user:=&model.User{}
	if err:=ctx.BindJSON(user);err!=nil{
		ctx.JSON(http.StatusOK,gin.H{
			"code": 10001,
			"msg":"error:"+err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":0,
		"msg":"",
		"data":user,
	})
}

func UserBatchStore(ctx *gin.Context)  {
	users:=&model.Users{}
	if err:=ctx.BindJSON(users);err!=nil {
		ctx.JSON(http.StatusOK,gin.H{
			"code": 10001,
			"msg":"error:"+err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":0,
		"msg":"",
		"data":users,
	})
}

func UserUpdate(ctx *gin.Context)  {
	ctx.String(http.StatusOK, "update")
}

func UserDestroy(ctx *gin.Context)  {
	ctx.String(http.StatusOK, "delete")
}