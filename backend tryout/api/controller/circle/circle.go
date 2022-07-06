package controller

import (
	"strconv"
	entity "studybuddy-backend-fast/api/entity"
	service "studybuddy-backend-fast/api/services/circle"

	"github.com/gin-gonic/gin"
)

type CircleController interface {
	FindAllCircle(c *gin.Context) []entity.Circle
	FindCircleByID(c *gin.Context) entity.Circle
	CreateCircle(c *gin.Context) entity.Circle
	UpdateCircle(c *gin.Context) entity.Circle
	DeleteCircle(c *gin.Context) error
}

type circleController struct {
	service service.CircleService
}

type deleteUserBody struct {
	id int
}

func NewCircleController(service service.CircleService) CircleController {
	return &circleController{service: service}
}

func (c *circleController) FindAllCircle(ctx *gin.Context) []entity.Circle {
	return c.service.FindAllCircle()
}

func (c *circleController) FindCircleByID(ctx *gin.Context) entity.Circle {
	circleID := ctx.Param("id")
	circle, _ := strconv.Atoi(circleID)
	return c.service.FindCircleByID(circle)

}

func (c *circleController) CreateCircle(ctx *gin.Context) entity.Circle {
	newCircle := entity.Circle{}
	if err := ctx.ShouldBind(&newCircle); err != nil {
		return newCircle
	}
	return c.service.CreateCircle(newCircle)
}

func (c *circleController) UpdateCircle(ctx *gin.Context) entity.Circle {
	newCircleData := entity.Circle{}
	if err := ctx.ShouldBind(&newCircleData); err != nil {
		return newCircleData
	}
	return c.service.UpdateCircle(newCircleData)
}

func (c *circleController) DeleteCircle(ctx *gin.Context) error {
	var circle entity.Circle
	err := ctx.ShouldBind(&circle)
	c.service.DeleteCircle(circle)
	if err == nil {
		return nil
	} else {
		return err
	}
}
