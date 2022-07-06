package routes

import (
	"github.com/gin-gonic/gin"
	controller "studybuddy-backend-fast/api/controller/school"
	auth "studybuddy-backend-fast/api/middlewares"
	repository "studybuddy-backend-fast/api/repository/school"
	service "studybuddy-backend-fast/api/services/school"
)

var (
	schoolRepo       repository.SchoolRepo       = repository.NewSchoolRepo()
	schoolService    service.SchoolService       = service.NewSchoolService(schoolRepo)
	schoolController controller.SchoolController = controller.NewSchoolController(schoolService)
)

func SchoolRoutes(route *gin.RouterGroup) {
	router := route.Group("/school")
	authMiddleware, _ := auth.New("admin-access")
	{
		router.GET("/", schoolController.FindSchool)
		router.GET("/:id", schoolController.FindSchoolByID)
	}

	authorized := router.Group("/school")
	authorized.Use(authMiddleware.MiddlewareFunc())
	{
		authorized.POST("/", schoolController.CreateSchool)
	}
}
