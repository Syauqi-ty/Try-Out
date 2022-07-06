package routes

import (
	controller "studybuddy-backend-fast/api/controller/student"
	auth "studybuddy-backend-fast/api/middlewares"
	repository "studybuddy-backend-fast/api/repository/student"
	testAuthRepository "studybuddy-backend-fast/api/repository/testauth"
	service "studybuddy-backend-fast/api/services/student"
	testAuthService "studybuddy-backend-fast/api/services/testauth"

	"github.com/gin-gonic/gin"
)

var (
	studentrepo       repository.StudentRepo          = repository.NewStudentRepo()
	tarepo            testAuthRepository.TestAuthRepo = testAuthRepository.NewTestAuthRepo()
	studentservice    service.StudentService          = service.NewStudentService(studentrepo)
	taservice         testAuthService.TestAuthService = testAuthService.NewTestAuthService(tarepo, studentrepo, testRepo)
	studentcontroller controller.StudentController    = controller.NewStudentController(studentservice, taservice)
)

func StudentRoutes(route *gin.RouterGroup) {
	router := route.Group("/student")
	authMiddleware, _ := auth.New("student-access")
	{
		router.POST("/auth", authMiddleware.LoginHandler)
		router.POST("/testauth", studentcontroller.TestAuth)
		router.POST("/newtestauth",studentcontroller.LoginSpaceBaru)
		router.POST("/", studentcontroller.CreateStudent)
		router.GET("/:id", studentcontroller.FindOneByID)
		router.PUT("/:id", studentcontroller.UpdateStudent)
		router.DELETE("/", studentcontroller.DeleteStudent)
		router.GET("/", studentcontroller.FindAll)
		router.POST("/forget",studentcontroller.Forget)
		router.Use(authMiddleware.MiddlewareFunc())
		{
			router.GET("/:id/testauth", studentcontroller.FindTestAuth)
		}
	}
}
