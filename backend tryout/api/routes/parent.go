package routes

import (
	"github.com/gin-gonic/gin"
	controller "studybuddy-backend-fast/api/controller/parent"
	auth "studybuddy-backend-fast/api/middlewares"
	repository "studybuddy-backend-fast/api/repository/parent"
	service "studybuddy-backend-fast/api/services/parent"
)

var (
	parentRepo       repository.ParentRepo       = repository.NewParentRepo()
	parentService    service.ParentService       = service.NewParentService(parentRepo)
	parentController controller.ParentController = controller.NewParentController(parentService)
)

func ParentRoutes(route *gin.RouterGroup) {
	router := route.Group("/parent")
	authMiddleware, _ := auth.New("admin-access")
	{
		router.POST("/login",parentController.Login)
		router.GET("/", parentController.FindAllParent)
		router.GET("/:id", parentController.FindParentByID)
	}

	// AUTHORIZED ROUTES
	authorizedRoutes := router.Group("/")
	authorizedRoutes.Use(authMiddleware.MiddlewareFunc())
	{
		authorizedRoutes.DELETE("/", parentController.DeleteParent)
		authorizedRoutes.POST("/", parentController.CreateParent)
	}
}
