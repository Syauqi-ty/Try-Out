package routes

import (
	"github.com/gin-gonic/gin"
	controller "studybuddy-backend-fast/api/controller/solution"
	auth "studybuddy-backend-fast/api/middlewares"
	repository "studybuddy-backend-fast/api/repository/solution"
	service "studybuddy-backend-fast/api/services/solution"
)

var (
	solutionRepo       repository.SolutionRepo       = repository.NewSolutionRepo()
	solutionService    service.SolutionService       = service.NewSolutionService(solutionRepo)
	solutionController controller.SolutionController = controller.NewSolutionController(solutionService)
)

func SolutionRoutes(route *gin.RouterGroup) {
	router := route.Group("/solution")
	authMiddleware, _ := auth.New("staff-access")
	{
		router.GET("/", solutionController.FindAllSolution)
		router.GET("/:id", solutionController.FindSolutionByID)
	}
	routerbaru := route.Group("/solutionbaru")
	routerbaru.GET("/:question_id",solutionController.FindSolutionByQID)
	router.Use(authMiddleware.MiddlewareFunc())
}
