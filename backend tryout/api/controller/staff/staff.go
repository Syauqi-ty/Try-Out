package controller

import (
	"strconv"
	entity "studybuddy-backend-fast/api/entity"
	shalala "studybuddy-backend-fast/api/helper"
	mapper "studybuddy-backend-fast/api/helper/querystring"
	service "studybuddy-backend-fast/api/services/staff"

	"github.com/gin-gonic/gin"
)

type StaffController interface {
	Save(ctx *gin.Context) error
	Update(ctx *gin.Context) error
	Delete(ctx *gin.Context) error
	FindAll(ctx *gin.Context)
	FindOneById(ctx *gin.Context) entity.Staff
	FindByAccess(ctx *gin.Context) []entity.Staff
	FindByAccessCheck(ctx *gin.Context) string
}

type staffController struct {
	service service.StaffService
}

func NewStaffController(service service.StaffService) StaffController {
	return &staffController{
		service: service,
	}
}

func (c *staffController) FindAll(ctx *gin.Context) {
	qstring, pagination := mapper.GetStaffMapper(ctx)
	res := c.service.FindAllStaff(qstring, pagination)
	pnum := strconv.Itoa(pagination["page"])
	lnum := strconv.Itoa(pagination["limit"])
	ctx.JSON(200, gin.H{"msg": "On page " + pnum + " with limit " + lnum, "data": res})
}

func (c *staffController) FindOneById(ctx *gin.Context) entity.Staff {
	hehe := ctx.Param("id")
	pantat, _ := strconv.Atoi(hehe)
	return c.service.FindById(pantat)
}

func (c *staffController) FindByAccess(ctx *gin.Context) []entity.Staff {
	return c.service.FindByAccess(ctx.Param("access"))
}

func (c *staffController) FindByAccessCheck(ctx *gin.Context) string {
	return c.service.FindByAccessCheck(ctx.Param("id"))
}

func (c *staffController) Save(ctx *gin.Context) error {
	var staff entity.Staff
	err := ctx.ShouldBindJSON(&staff)
	if err != nil {
		return err
	}
	staff.Password = shalala.Encrypt(staff.Password)
	c.service.Save(staff)
	return nil
}

func (c *staffController) Update(ctx *gin.Context) error {
	var staff entity.Staff
	err := ctx.ShouldBindJSON(&staff)
	c.service.Update(staff)
	if err == nil {
		return nil
	} else {
		return err
	}
}

func (c *staffController) Delete(ctx *gin.Context) error {
	var staff entity.Staff
	err := ctx.ShouldBindJSON(&staff)
	c.service.Delete(staff)
	if err == nil {
		return nil
	} else {
		return err
	}
}
