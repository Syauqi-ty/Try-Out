package controller

import (
	entity "studybuddy-backend-fast/api/entity"
	service "studybuddy-backend-fast/api/services/feedback"
	"github.com/gin-gonic/gin"
)

type FeedBackController interface {
	CreateFeedBack(c *gin.Context)
	FindAllFeedback(c *gin.Context)
}

type feedbackController struct {
	service service.FeedbackService
}

func NewFeedbackController(service service.FeedbackService) FeedBackController {
	return &feedbackController{service}
}

func (s *feedbackController) CreateFeedBack(c *gin.Context) {
	var feedback entity.Feedback
	if err := c.ShouldBind(&feedback); err == nil {
		res := s.service.CreateFeedback(feedback)
		c.JSON(200, gin.H{"msg": "Succesfully Posted", "data": res})
	} else {
		c.JSON(400, gin.H{"msg": err.Error()})
	}
}
 
func (s *feedbackController) FindAllFeedback(c *gin.Context) {
	data := s.service.FindAllFeedback()
	c.JSON(200, gin.H{"data":data})
}