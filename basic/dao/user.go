package dao

import (
	"juggle/basic/db"
	"juggle/basic/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UserList(ctx *gin.Context) {
	query := &model.UserQuery{}
	if err := ctx.BindQuery(query); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 10000,
			"msg":  "error:" + err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "",
		"data": query,
	})
}

func UserShow(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "",
		"data": model.NewUser(uint(id), "张三"),
	})
}

func UserStore(ctx *gin.Context) {
	user := &model.User{}
	if err := ctx.BindJSON(user); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 10001,
			"msg":  "error:" + err.Error(),
		})
		return
	}

	result := db.DB().Create(user)
	if result.Error != nil || result.RowsAffected == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 10002,
			"msg":  result.Error.Error(),
			"data": nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "",
		"data": user.ID,
	})
}

func UserBatchStore(ctx *gin.Context) {
	users := &model.Users{}
	if err := ctx.BindJSON(users); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 10001,
			"msg":  "error:" + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "",
		"data": users,
	})
}

func UserUpdate(ctx *gin.Context) {
	ctx.String(http.StatusOK, "update")
}

func UserDestroy(ctx *gin.Context) {
	ctx.String(http.StatusOK, "delete")
}
