package controller

import (
	"villa-akmali/api/model"
	service "villa-akmali/api/service/user"

	"github.com/gin-gonic/gin"
)


type UserController interface {
	CreateUser(ctx *gin.Context)
	Login(ctx *gin.Context)
}

type userController struct {
	service service.UserService
} 

type DataReturn struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Level int `json:"level"`
}

func NewUserController(service service.UserService) UserController{
	  return &userController{service: service}
}

func (c *userController) CreateUser(ctx *gin.Context)  {
	var user model.User
	var datauser DataReturn
	ctx.ShouldBind(&user)
	data := c.service.CreateUser(user)
	if data.ID == 0 {
		ctx.JSON(400,gin.H{"msg":"user sudah terdaftar"})
	} else {
		datauser.ID = data.ID
		datauser.Name = data.Name
		datauser.Email = data.Email
		datauser.Level = data.Level
		datauser.Phone = data.Phone
		ctx.JSON(200,gin.H{"data": datauser })
	}
}

func (c *userController) Login(ctx *gin.Context)  {
	var user model.Login
	var datauser DataReturn
	ctx.ShouldBind(&user)
	data := c.service.Login(user)
	if data.Name== "" {
		ctx.JSON(400,gin.H{"msg":"Ada yang Salah"})
	} else if data.Name == "123456789" && data.ID == 0 {
		ctx.JSON(404,gin.H{"msg":"password salah"})
	} else {
		datauser.ID = data.ID
		datauser.Name = data.Name
		datauser.Email = data.Email
		datauser.Phone = data.Phone
		datauser.Level = data.Level
		ctx.JSON(200,gin.H{"data": datauser })
	}
}