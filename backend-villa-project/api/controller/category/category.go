package controller

import (
	"villa-akmali/api/model"
	service "villa-akmali/api/service/category"

	"github.com/gin-gonic/gin"
)


type CategoryController interface {
	CreateCategory(ctx *gin.Context)
	FindAll(ctx *gin.Context)
}

type categoryController struct {
	service service.CategoryService
}

func NewCategoryController(service service.CategoryService) CategoryController {
	return &categoryController{
		service : service,
	}
}

func (c *categoryController) CreateCategory(ctx *gin.Context) {
	var category model.Category
	err := ctx.ShouldBind(&category)
	if err != nil {
		ctx.JSON(400,gin.H{"msg":"jangan kosong dong"})
	}else{
		data := c.service.CreateCategory(category)
		if data.ID == 0 {
			ctx.JSON(400,gin.H{"msg":"jangan kosong dong"})
		}else{
			ctx.JSON(200,gin.H{"data":data})
		}
	}
}

func (c *categoryController) FindAll(ctx *gin.Context)  {
	data := c.service.FindAll()
	ctx.JSON(200,gin.H{"data":data})
}
