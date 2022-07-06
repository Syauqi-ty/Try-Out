package controller

import (
	"strconv"
	entity "studybuddy-backend-fast/api/entity"
	mapper "studybuddy-backend-fast/api/helper/querystring"
	service "studybuddy-backend-fast/api/services/parent"

	"github.com/gin-gonic/gin"
)

type ParentController interface {
	FindAllParent(c *gin.Context)
	FindParentByID(c *gin.Context)
	CreateParent(c *gin.Context)
	UpdateParent(c *gin.Context)
	DeleteParent(c *gin.Context)
	Login(c *gin.Context)
}

type parentController struct {
	service service.ParentService
}

func NewParentController(service service.ParentService) ParentController {
	return &parentController{service}
}

func (s *parentController) FindAllParent(c *gin.Context) {
	qstring, pag := mapper.GetParentMapper(c)
	pagenum := strconv.Itoa(pag["page"])
	limnum := strconv.Itoa(pag["limit"])

	res := s.service.FindAllParent(qstring, pag)
	c.JSON(200, gin.H{"msg": "On page " + pagenum + " with limit " + limnum, "data": res})
}

func (s *parentController) FindParentByID(c *gin.Context) {
	if id, err := strconv.Atoi(c.Param("id")); err == nil {
		res := s.service.FindParentByID(id)
		c.JSON(200, gin.H{"msg": "Query successful", "data": res})
	} else {
		c.JSON(400, gin.H{"msg": "Bad request, invalid id"})
	}
}

func (s *parentController) CreateParent(c *gin.Context) {
	var parent entity.Parent
	if err := c.ShouldBind(&parent); err == nil {
		res := s.service.CreateParent(parent)
		c.JSON(200, gin.H{"msg": "Parent created", "data": res})
	} else {
		c.JSON(400, gin.H{"msg": err.Error()})
	}
}

func (s *parentController) UpdateParent(c *gin.Context) {
	var parent entity.Parent
	if err := c.ShouldBind(&parent); err == nil {
		s.service.UpdateParent(parent)
		c.JSON(200, gin.H{"msg": "Parent updated"})
	} else {
		c.JSON(400, gin.H{"msg": err.Error()})
	}
}
func (s *parentController) Login(c *gin.Context){
	var parent entity.Parent
	if err := c.ShouldBindJSON(&parent); err == nil {
		data:=s.service.Login(parent)
		if data.ID != 0 {
			c.JSON(200, gin.H{"data": data})
		}else if data.ID == 0{
			c.JSON(400, gin.H{"data": nil})
		}
	} else {
		c.JSON(400, gin.H{"msg": err.Error()})
	}
	
} 

func (s *parentController) DeleteParent(c *gin.Context) {
	var parent entity.Parent
	if err := c.ShouldBind(&parent); err == nil {
		s.service.DeleteParent(parent)
		c.JSON(200, gin.H{"msg": "Parent deleted"})
	} else {
		c.JSON(400, gin.H{"msg": err.Error()})
	}
}