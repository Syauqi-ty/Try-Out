package controller

import (
	"github.com/gin-gonic/gin"
	"strconv"
	entity "studybuddy-backend-fast/api/entity"
	service "studybuddy-backend-fast/api/services/operator"
	mapper "studybuddy-backend-fast/api/helper/querystring"
)

type OperatorController interface{
	FindOperator(c *gin.Context)
	FindOperatorByID(c *gin.Context)
	CreateOperator(c *gin.Context)
	UpdateOperator(c *gin.Context)
	DeleteOperator(c *gin.Context)
}

type operatorController struct {
	service service.OperatorService
}

func NewOperatorController(service service.OperatorService) OperatorController {
	return &operatorController{service}
}

func (s *operatorController) FindOperator(c *gin.Context) {
	param, pagination := mapper.GetParentMapper(c)
	pagenum := strconv.Itoa(pagination["page"])
	limnum := strconv.Itoa(pagination["limit"])

	res := s.service.FindOperator(param, pagination)
	c.JSON(200, gin.H{"msg": "On page " + pagenum + " with limit " + limnum, "data": res})
}

func (s *operatorController) FindOperatorByID(c *gin.Context) {
	if id, err := strconv.Atoi(c.Param("id")); err == nil {
		res := s.service.FindOperatorByID(id)
		c.JSON(200, gin.H{"msg": "Query successful", "data": res})
	} else {
		c.JSON(400, gin.H{"msg": err.Error()})
	}
}

func (s *operatorController) CreateOperator(c *gin.Context) {
	var operator entity.Operator
	if err := c.ShouldBind(&operator); err == nil {
		res := s.service.CreateOperator(operator)
		c.JSON(200, gin.H{"msg": "Operator created", "data": res})
	} else {
		c.JSON(400, gin.H{"msg": err.Error()})
	}
}

func (s *operatorController) UpdateOperator(c *gin.Context) {
	var operator entity.Operator
	if err := c.ShouldBind(&operator); err == nil {
		s.service.UpdateOperator(operator)
		c.JSON(200, gin.H{"msg": "Operator updated"})
	} else {
		c.JSON(400, gin.H{"msg": err.Error()})
	}
}

func (s *operatorController) DeleteOperator(c *gin.Context) {
	var operator entity.Operator
	if err := c.ShouldBind(&operator); err == nil {
		s.service.DeleteOperator(operator)
		c.JSON(200, gin.H{"msg": "Operator deleted"})
	} else {
		c.JSON(400, gin.H{"msg": err.Error()})
	}
}