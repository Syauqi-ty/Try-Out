package controller

import (
	"github.com/gin-gonic/gin"
	"strconv"
	entity "studybuddy-backend-fast/api/entity"
	service "studybuddy-backend-fast/api/services/question"
)

type QuestionController interface {
	CreateQuestion(c *gin.Context) entity.Question
	FindAllQuestion(c *gin.Context)
	FindQuestionByID(c *gin.Context) entity.Question
	UpdateQuestion(c *gin.Context) int
	DeleteQuestion(c *gin.Context) int
	TemporaryDelete(c *gin.Context) int
	QuestionandSolution(c *gin.Context)
}

type questionController struct {
	service service.QuestionService
}
type Body struct {
	Soal entity.Question `json:"soal"`
	Solusi entity.Solution `json:"solusi"`
}

func NewQuestionController(s service.QuestionService) QuestionController {
	return &questionController{service: s}
}

func (s *questionController) CreateQuestion(c *gin.Context) entity.Question {
	var question entity.Question
	if err := c.ShouldBind(&question); err != nil {
		return question
	}
	return s.service.CreateQuestion(question)
}

func (s *questionController) FindAllQuestion(c *gin.Context) {
	res := s.service.FindAllQuestionMin()
	c.JSON(200, gin.H{"msg": "Questions fetched", "data": res})
}

func (s *questionController) FindQuestionByID(c *gin.Context) entity.Question {
	id, _ := strconv.Atoi(c.Param("id"))
	return s.service.FindQuestionByID(id)
	// nanti benerin
	// if err != nil {
	// 	return "PATNAT"
	// }
}

func (s *questionController) UpdateQuestion(c *gin.Context) int {
	var question entity.Question
	if err := c.ShouldBind(&question); err != nil {
		return 0
	}
	data := s.service.UpdateQuestion(question)
	if data == 0 {
		return 0
	} else {
		return 1
	}
}

func (s *questionController) DeleteQuestion(c *gin.Context) int {
	id, _ := strconv.Atoi(c.Param("id"))
	if id == 0 {
		return 0
	} else{
		s.service.DeleteQuestion(id)
		return 1
	}
}
func (s *questionController) TemporaryDelete(c *gin.Context) int {
	id, _ := strconv.Atoi(c.Param("id"))
	if id == 0 {
		return 0
	} else{
		s.service.TemporaryDelete(id)
		return 1
	}
}
func (s *questionController) QuestionandSolution(c *gin.Context) {
	var body Body
	bodyjson := c.ShouldBind(&body)
	if bodyjson == nil {
		c.JSON(200,gin.H{"status":"data"})
	} else{
		c.JSON(400, gin.H{"msg": bodyjson.Error()})
	}

}
