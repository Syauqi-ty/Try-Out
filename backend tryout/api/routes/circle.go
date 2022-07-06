package routes

import (
	controller "studybuddy-backend-fast/api/controller/circle"
	repository "studybuddy-backend-fast/api/repository/circle"
	service "studybuddy-backend-fast/api/services/circle"

	"github.com/gin-gonic/gin"
)

var (
	circleRepo       repository.CircleRepo       = repository.NewCircleRepo()
	circleService    service.CircleService       = service.NewCircleService(circleRepo)
	circleController controller.CircleController = controller.NewCircleController(circleService)
)

func findAllCircle(c *gin.Context) {
	res := circleController.FindAllCircle(c)
	c.JSON(200, gin.H{"msg": "Query successful", "data": res})
}

func findCircleByID(c *gin.Context) {
	res := circleController.FindCircleByID(c)
	c.JSON(200, gin.H{"msg": "Query successful", "data": res})
}

func createCircle(c *gin.Context) {
	res := circleController.CreateCircle(c)
	c.JSON(200, gin.H{"msg": "Query successful", "data": res})
}

func updateCircle(c *gin.Context) {
	res := circleController.UpdateCircle(c)
	c.JSON(200, gin.H{"msg": "Query successful", "data": res})
}

func CircleRoutes(router *gin.RouterGroup) {
	route := router.Group("/circle")

	{
		route.GET("/", findAllCircle)
		route.GET("/:id", findCircleByID)
		route.PUT("/:id", updateCircle)
		route.POST("/", createCircle)
	}
}
