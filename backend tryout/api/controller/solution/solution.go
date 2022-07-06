package controller

import (
	"github.com/gin-gonic/gin"
	"strconv"
	service "studybuddy-backend-fast/api/services/solution"
)

type SolutionController interface {
	FindAllSolution(c *gin.Context)
	FindSolutionByID(c *gin.Context)
	FindSolutionByQID(c *gin.Context)
}

type solutionController struct {
	service service.SolutionService
}

func NewSolutionController(service service.SolutionService) SolutionController {
	return &solutionController{service}
}

func (s *solutionController) FindAllSolution(c *gin.Context) {
	res := s.service.FindAllSolution()
	c.JSON(200, gin.H{"msg": "Query successful", "data": res})
}

func (s *solutionController) FindSolutionByID(c *gin.Context) {
	if id, err := strconv.Atoi(c.Param("id")); err == nil {
		res := s.service.FindSolutionByID(id)
		c.JSON(200, gin.H{"msg": "Query successful", "data": res})
	} else {
		c.JSON(400, gin.H{"msg": "Bad request, nothing fetched"})
	}
}
func (s *solutionController) FindSolutionByQID(c *gin.Context){
	if id, err := strconv.Atoi(c.Param("question_id")); err == nil {
		res := s.service.FindSolutionByQID(id)
		c.JSON(200, gin.H{"msg": "Query successful", "data": res})
	} else {
		c.JSON(400, gin.H{"msg": "Bad request, nothing fetched"})
	}
}