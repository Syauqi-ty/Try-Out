package routes

import (
	controller "studybuddy-backend-fast/api/controller/feedback"
	repository "studybuddy-backend-fast/api/repository/feedback"
	service "studybuddy-backend-fast/api/services/feedback"

	"github.com/gin-gonic/gin"
)

var (
	feedbackrepo repository.FeedbackRepo = repository.NewFeedbackRepo()
	feedbackservice service.FeedbackService = service.NewFeedbackService(feedbackrepo)
	feedbackcontroller controller.FeedBackController = controller.NewFeedbackController(feedbackservice)
)

func FeedbackRoutes(route *gin.RouterGroup){
	router := route.Group("/feedback")
	{
		router.GET("/",feedbackcontroller.FindAllFeedback)
		router.POST("/create",feedbackcontroller.CreateFeedBack)
	}
}