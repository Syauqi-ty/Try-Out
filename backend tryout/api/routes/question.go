package routes

import (
	"github.com/gin-gonic/gin"
	controller "studybuddy-backend-fast/api/controller/question"
	auth "studybuddy-backend-fast/api/middlewares"
	repository "studybuddy-backend-fast/api/repository/question"
	service "studybuddy-backend-fast/api/services/question"
)

var (
	questionRepo       repository.QuestionRepo       = repository.NewQuestionRepo()
	questionService    service.QuestionService       = service.NewQuestionService(questionRepo)
	questionController controller.QuestionController = controller.NewQuestionController(questionService)
)

func findQuestionByID(c *gin.Context) {
	res := questionController.FindQuestionByID(c)
	c.JSON(200, gin.H{"msg": "Query successful", "data": res})
}

func createQuestion(c *gin.Context) {
	res := questionController.CreateQuestion(c)
	c.JSON(200, gin.H{"msg": "Query successful", "data": res})
}

func updateQuestion(c *gin.Context) {
	res := questionController.UpdateQuestion(c)
	if res == 0 {
		c.JSON(400, gin.H{"error": "Something went wrong"})
	} else {
		c.JSON(200, gin.H{"messages": "Hamdabi Jaya Jaya Jaya"})
	}
}

func deleteQuestion(c *gin.Context) {
	err := questionController.DeleteQuestion(c)
	if err == 0 {
		c.JSON(400, gin.H{"msg": "Bad request, nothing deleted"})
	} else {
		c.JSON(200, gin.H{"msg": "Question deleted"})
	}
}
func temporaryDelete(c *gin.Context) {
	err := questionController.TemporaryDelete(c)
	if err == 0 {
		c.JSON(400, gin.H{"msg": "Bad request, nothing deleted"})
	} else {
		c.JSON(200, gin.H{"msg": "Question updated"})
	}
}

func QuestionRoutes(router *gin.RouterGroup) {
	route := router.Group("/question")
	authMiddleware, _ := auth.New("staff-access")
	{
		route.GET("/", questionController.FindAllQuestion)
		route.GET("/:id", findQuestionByID)
		route.PUT("/tempo/:id", temporaryDelete)

		staffRouter := route.Group("/")
		staffRouter.Use(authMiddleware.MiddlewareFunc())
		{
			staffRouter.PUT("/", updateQuestion)
			staffRouter.POST("/", createQuestion)
			staffRouter.DELETE("/:id", deleteQuestion)
		}
	}
}
