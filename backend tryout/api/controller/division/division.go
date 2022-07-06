package controller

import (
	"strconv"
	entity "studybuddy-backend-fast/api/entity"
	service "studybuddy-backend-fast/api/services/division"

	"github.com/gin-gonic/gin"
)
type DivisionController interface {
	CreateDivision(c *gin.Context)
	FindAll(c *gin.Context)
	Description(c *gin.Context)
}

type divisionController struct {
	service service.DivisionService
}
func NewDivisionController(service service.DivisionService) DivisionController {
	return &divisionController{service: service}
}

func (s *divisionController) CreateDivision(c *gin.Context) {
	var divisi entity.Division
	if err := c.ShouldBind(&divisi); err == nil {
		res := s.service.CreateDivision(divisi)
		c.JSON(200, gin.H{"msg": "Succesfully Posted", "data": res})
	} else {
		c.JSON(400, gin.H{"msg": err.Error()})
	}
}
func (s *divisionController) FindAll(c *gin.Context) {
	data := s.service.FindAllDivisi()
	c.JSON(200,gin.H{"data":data})
}
func (s *divisionController) Description(c *gin.Context) {
	divid,_ := strconv.Atoi(c.Param("divid"))
	posid,_ := strconv.Atoi(c.Param("posid"))
	data := s.service.Description(divid,posid)
	c.JSON(200,gin.H{"data":data})
}