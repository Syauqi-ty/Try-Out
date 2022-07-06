package routes

import (
	"github.com/gin-gonic/gin"
	controller "studybuddy-backend-fast/api/controller/pref"
	auth "studybuddy-backend-fast/api/middlewares"
	repository "studybuddy-backend-fast/api/repository/pref"
	service "studybuddy-backend-fast/api/services/pref"
)

var (
	prefRepo       repository.PrefRepo       = repository.NewPrefRepo()
	prefService    service.PrefService       = service.NewPrefService(prefRepo)
	prefController controller.PrefController = controller.NewPrefController(prefService, taservice)
)

func PrefRoutes(route *gin.RouterGroup) {
	router := route.Group("/pref")
	authMiddleware, _ := auth.New("student-access")
	{
		router.GET("/", prefController.FindAllUni)
		router.GET("/:id", prefController.FindUniByID)
	}

	studentRouter := router.Group("/")
	studentRouter.Use(authMiddleware.MiddlewareFunc())
	{
		studentRouter.PUT("/add", prefController.SetPref)
	}
}
