package controller

import (
	"github.com/gin-gonic/gin"
	"strconv"
	entity "studybuddy-backend-fast/api/entity"
	service "studybuddy-backend-fast/api/services/school"
)

type SchoolController interface {
	FindSchool(c *gin.Context)
	FindSchoolByID(c *gin.Context)
	CreateSchool(c *gin.Context)
	UpdateSchool(c *gin.Context)
	DeleteSchool(c *gin.Context)
}

type schoolController struct {
	service service.SchoolService
}

func NewSchoolController(s service.SchoolService) SchoolController {
	return &schoolController{s}
}

func (s *schoolController) FindSchool(c *gin.Context) {
	if slug := c.DefaultQuery("slug", ""); slug != "" {
		res := s.service.FindSchoolBySlug(slug)
		c.JSON(200, gin.H{"msg": "Query successful", "data": res})
	} else {
		res := s.service.FindAllSchool()
		c.JSON(200, gin.H{"msg": "Query successul", "data": res})
	}
}

func (s *schoolController) FindSchoolByID(c *gin.Context) {
	if id, err := strconv.Atoi(c.Param("id")); err == nil {
		res := s.service.FindSchoolByID(id)
		c.JSON(200, gin.H{"msg": "Query successful", "data": res})
	} else {
		c.JSON(400, gin.H{"msg": err.Error()})
	}
}

func (s *schoolController) CreateSchool(c *gin.Context) {
	var school entity.School
	if err := c.ShouldBind(&school); err == nil {
		s.service.CreateSchool(school)
		c.JSON(200, gin.H{"msg": "School created", "data": school})
	} else {
		c.JSON(400, gin.H{"msg": err.Error()})
	}
}

func (s *schoolController) UpdateSchool(c *gin.Context) {
	var school entity.School
	if err := c.ShouldBind(&school); err == nil {
		s.service.UpdateSchool(school)
		c.JSON(200, gin.H{"msg": "School data updated", "data": school})
	} else {
		c.JSON(400, gin.H{"msg": err.Error()})
	}
}

func (s *schoolController) DeleteSchool(c *gin.Context) {
	var school entity.School
	if err := c.ShouldBind(&school); err == nil {
		s.service.DeleteSchool(school)
		c.JSON(200, gin.H{"msg": "School deleted"})
	} else {
		c.JSON(400, gin.H{"msg": err.Error()})
	}
}
