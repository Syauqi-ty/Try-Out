package api

import (
	"studybuddy-backend-fast/api/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/hipcar/go-xendit-client/xendit"
	"github.com/spf13/viper"
)

func setupRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(gin.Logger())

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AddAllowHeaders("Access-Control-Allow-Origin")
	config.AddAllowHeaders("Authorization")

	router.Use(cors.New(config))

	router.LoadHTMLGlob("api/templates/*")

	v2 := router.Group("/api/v2")
	{
		routes.TestRoutes(v2)
		routes.StudentRoutes(v2)
		routes.AuthRoutes(v2)
		routes.ApplicantRoutes(v2)
		routes.CircleRoutes(v2)
		routes.DisbursementRoutes(v2)
		routes.StaffRoutes(v2)
		routes.QuestionRoutes(v2)
		routes.SolutionRoutes(v2)
		routes.ScoreRoutes(v2)
		routes.PrefRoutes(v2)
		routes.PaymentRoutes(v2)
		routes.FileRoutes(v2)
		routes.SchoolRoutes(v2)
		routes.ParentRoutes(v2)
		routes.FeedbackRoutes(v2)
		routes.DivisionRoutes(v2)
		routes.PosisiRoutes(v2)
		routes.RecruitmentRoutes(v2)
	}

	return router
}

func initVendors() {
	// Xendit Client Init
	xenditClient := xendit.NewClient()
	xenditClient.EnableLog = true // logging is false by default
	xenditClient.SecretKey = "xnd_development_DUZflxQtZ9sBgILME2ooOLgHhvgoK9xDUnV3RTWSE1jd7NOel8QvfSyWOq0OUOv"
}

func Run() {
	router := setupRouter()
	router.Run(viper.GetString(`server.address`))
}
