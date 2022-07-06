package routes

import (
	controller "studybuddy-backend-fast/api/controller/score"
	auth "studybuddy-backend-fast/api/middlewares"
	repository "studybuddy-backend-fast/api/repository/score"
	studentRepository "studybuddy-backend-fast/api/repository/student"
	testRepository "studybuddy-backend-fast/api/repository/test"
	testAuthRepository "studybuddy-backend-fast/api/repository/testauth"
	service "studybuddy-backend-fast/api/services/score"

	"github.com/gin-gonic/gin"
)

var (
	scoreRepo       repository.ScoreRepo            = repository.NewScoreRepo()
	testRepo        testRepository.TestRepo         = testRepository.NewTestRepo()
	testAuthRepo    testAuthRepository.TestAuthRepo = testAuthRepository.NewTestAuthRepo()
	studentRepo     studentRepository.StudentRepo   = studentRepository.NewStudentRepo()
	scoreService    service.ScoreService            = service.NewScoreService(scoreRepo, testRepo, studentRepo, questionRepo, testAuthRepo)
	scoreController controller.ScoreController      = controller.NewScoreController(scoreService)
)

func findAllScore(c *gin.Context) {
	res := scoreController.FindAllScore(c)
	c.JSON(200, gin.H{"msg": "Query successful", "data": res})
}

func ScoreRoutes(route *gin.RouterGroup) {
	router := route.Group("/score")
	authMiddleware, _ := auth.New("student-access")
	{
		router.GET("/", findAllScore)
		router.GET("/:id", scoreController.AllStudentScore)
		router.GET("/:id/highscore", scoreController.HighScore)
		router.GET("/:id/battlehigh", scoreController.HighScoreBattle)
		router.GET("/:id/graph", scoreController.GraphicData)
		router.GET("/:id/stat", scoreController.UserStats)
		router.GET("/:id/battle", scoreController.Battled)
		router.GET("/:id/avg", scoreController.AvgTKA)
		router.GET("/:id/quartile", scoreController.Quartile)
		router.GET("/:id/rara", scoreController.Rara)
		router.GET("/:id/distribution", scoreController.Distribution)
		router.GET("/:id/leaderboard", scoreController.LeaderBoard)
		router.GET("/:id/stat/test", scoreController.TestStats)
		router.GET("/:id/rank/:studentID", scoreController.GetRank)
		router.GET("/:id/true/:testid",scoreController.BenarSalah)
		router.GET("/:id/newleaderboard/:studentID",scoreController.NewLeaderBoard)

		// Authorized routes
		studentRouter := router.Group("/")
		studentRouter.Use(authMiddleware.MiddlewareFunc())
		{
			studentRouter.GET("/:id/check", scoreController.CheckTest)
			studentRouter.POST("/jawab",scoreController.Jawab)
		}
		customRoutes := route.Group("/nilai")
		{
			customRoutes.GET("/", scoreController.Graphic)
		}
	}
}
