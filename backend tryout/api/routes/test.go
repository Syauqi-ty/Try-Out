package routes

import (
	controller "studybuddy-backend-fast/api/controller/test"
	auth "studybuddy-backend-fast/api/middlewares"
	repository "studybuddy-backend-fast/api/repository/test"
	service "studybuddy-backend-fast/api/services/test"

	"github.com/gin-gonic/gin"
)

var (
	testrepo       repository.TestRepo       = repository.NewTestRepo()
	testservice    service.TestService       = service.NewTestService(testrepo, scoreRepo)
	testcontroller controller.TestController = controller.NewTestController(testservice, scoreService)
)

func findTest(ctx *gin.Context) {
	res := testcontroller.FindTestMin(ctx)
	ctx.JSON(200, gin.H{"msg": "Query successful", "data": res})
}

func createTest(ctx *gin.Context) {
	data := testcontroller.CreateTest(ctx)
	ctx.JSON(200, gin.H{"msg": "New test created", "data": data})
}
func updateTest(ctx *gin.Context) {
	res := testcontroller.UpdateTest(ctx)
	ctx.JSON(200, gin.H{"msg": "Test updated", "data": res})
}

func deleteTest(ctx *gin.Context) {
	err := testcontroller.DeleteTest(ctx)
	if err == nil {
		ctx.JSON(200, gin.H{"msg": "Test deleted"})
	} else {
		ctx.JSON(400, gin.H{"msg": "Bad request, nothing was deleted"})
	}
}

func findTestWithQuestions(c *gin.Context) {
	res := testcontroller.FindTest(c)
	c.JSON(200, gin.H{"msg": "Query successful", "data": res})
}

func TestRoutes(route *gin.RouterGroup) {
	router := route.Group("/test")
	anotherAuthMiddleware, _ := auth.New("student-access")
	authMiddleware, _ := auth.New("admin-access")
	{
		router.GET("/", testcontroller.FindAllTest)
		router.GET("/:id", findTest)
		router.GET("/:id/battled", testcontroller.IkutTest)
		router.GET("/:id/soal/:subtest",testcontroller.SoalBattle)

	}

	// Authorized routes
	studentRouter := router.Group("/")
	studentRouter.Use(anotherAuthMiddleware.MiddlewareFunc())
	{
		studentRouter.GET("/:id/all", findTestWithQuestions)
		studentRouter.GET("/:id/all/solution", testcontroller.FindTestFull)
		studentRouter.PUT("/:testID/answer/:questionID", testcontroller.AnswerQuestion)
	}

	adminRouter := router.Group("/")
	adminRouter.Use(authMiddleware.MiddlewareFunc())
	{
		adminRouter.POST("/", createTest)
		adminRouter.PUT("/", updateTest)
		adminRouter.DELETE("/", deleteTest)
	}

	customRouter := route.Group("/custom_test")
	{
		customRouter.GET("/last", testcontroller.LastBattle)
		router.GET("/:id/available", testcontroller.AvailableTest)
	}
}
