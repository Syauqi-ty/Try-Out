package controller

import (
	"strconv"
	"villa-akmali/api/helper"
	"villa-akmali/api/model"
	service "villa-akmali/api/service/order"

	"github.com/gin-gonic/gin"
)


type OrderController interface {
	CreateOrder(ctx *gin.Context)
	FindAll(ctx *gin.Context)
	Update(ctx *gin.Context)
	Yearly(ctx *gin.Context) 
	Monthly(ctx *gin.Context)
	Weekly(ctx *gin.Context)
	GraphYearly(ctx *gin.Context)
	GraphMonth(ctx *gin.Context)
	GraphWeek(ctx *gin.Context)
	SumAll(ctx *gin.Context)
	FindByID(ctx *gin.Context)
	Filter(ctx *gin.Context) 
}

type orderController struct {
	service service.OrderService
}

func NewOrderService(service service.OrderService) OrderController {
	return &orderController{service: service,}
}

func (c *orderController) CreateOrder(ctx *gin.Context)  {
	var order model.Order
	err := ctx.ShouldBind(&order)
	if err != nil {
		ctx.JSON(400,gin.H{"msg":"inputan yang bener mas"})
	} else {
		data := c.service.CreateOrder(order)
		if data.CategoryID == 0{
			ctx.JSON(400,gin.H{"msg":"inputtan kategorinya yang bener"})
		} else {
			ctx.JSON(200,gin.H{"data":data})
		}
	}
}

func (c *orderController) FindAll(ctx *gin.Context) {
	pagination := helper.Pagination(ctx)
	data,page := c.service.AllOrder(pagination)
	ctx.JSON(200,gin.H{"data":data,"limit":page})
}

func (c *orderController) FindByID(ctx *gin.Context) {
	id,_ := strconv.Atoi(ctx.Param("id"))
	data := c.service.FindById(id)
	if data.ID == 0{
		ctx.JSON(400,gin.H{"msg":"gaada data ges"})
	}else{
		ctx.JSON(200,gin.H{"data":data})
	}
}

func (c *orderController) Update(ctx *gin.Context)  {
	var order model.Order
	id,_ := strconv.Atoi(ctx.Param("id")) 
	err := ctx.ShouldBind(&order)
	if err != nil {
		ctx.JSON(400,gin.H{"msg" : "Yang bener ges"})
	} else {
		c.service.Update(id,order)
		ctx.JSON(200,gin.H{"msg":"nice one ges"})
	}
}

func (c *orderController) Yearly(ctx *gin.Context)  {
	data := c.service.Yearly()
	ctx.JSON(200,gin.H{"data":data})
}
func (c *orderController) GraphYearly(ctx *gin.Context)  {
	data := c.service.GraphYear()
	ctx.JSON(200,gin.H{"data":data})
}
func (c *orderController) Monthly(ctx *gin.Context)  {
	data := c.service.Monthly()
	ctx.JSON(200,gin.H{"data":data})
}

func (c *orderController) Weekly(ctx *gin.Context)  {
	data := c.service.Weekly()
	ctx.JSON(200,gin.H{"data":data})
}
func (c *orderController) SumAll(ctx *gin.Context)  {
	data := c.service.SumAll()
	ctx.JSON(200,gin.H{"data":data})
}

func (c *orderController) GraphMonth(ctx *gin.Context)  {
	data := c.service.GraphMonth()
	ctx.JSON(200,gin.H{"data":data})
}

func (c *orderController) GraphWeek(ctx *gin.Context)  {
	data := c.service.GraphWeek()
	ctx.JSON(200,gin.H{"data":data})
}

func (c *orderController) Filter(ctx *gin.Context) {
	var filter model.QueryFilter
	err := ctx.ShouldBind(&filter)
	if err != nil{
		ctx.JSON(400,gin.H{"msg":"yang bener aja ah"})
	}else{
		data := c.service.Filter(filter)
		ctx.JSON(200,gin.H{"data":data})
	}
}