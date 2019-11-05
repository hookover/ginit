package controllers

import (
	"gin_api/app/models"
	"gin_api/database/db"
	"gin_api/database/redis"
	"gin_api/util/response"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"time"
)

type LoginRequest struct {
	Username string `binding:"required" json:"username"`
	Password string `binding:"required" json:"password"`
}

func Login(ctx *gin.Context) {
	username := ctx.DefaultQuery("username", "")
	password := ctx.DefaultQuery("password", "")
	utilGin := response.Gin{Ctx: ctx}

	if err := ctx.ShouldBindWith(&LoginRequest{Username: username, Password: password}, binding.Query); err != nil {
		utilGin.Response(400, err.Error(), nil)
		return
	}
	redis.Chan("local").Set("test", username, time.Minute*10)

	users := make([]models.User, 0)

	err := db.Chan("local").Find(&users)
	if err != nil {
		utilGin.Response(500, err.Error(), nil)
	} else {
		utilGin.Response(200, "success", gin.H{
			"cache":    redis.Chan("local").Get("test").Val(),
			"username": username, "password": password, "users": users})
	}
}

func Register(ctx *gin.Context) {
	email := ctx.DefaultQuery("email", "")
	phone := ctx.DefaultQuery("phone", "")
	username := ctx.DefaultQuery("username", "")
	password := ctx.DefaultQuery("password", "")

	user := &models.User{
		Name:     username,
		Password: password,
		Email:    email,
		Phone:    phone,
	}

	id, err := db.Chan("local").Insert(user)

	if err != nil {
		ctx.JSON(500, gin.H{"status": 500, "error": err.Error()})
	} else {
		ctx.JSON(200, gin.H{"status": 200, "data": id})
	}
}
