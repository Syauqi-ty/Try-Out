package routes

import (
	controller "studybuddy-backend-fast/api/controller/applicant"
	auth "studybuddy-backend-fast/api/middlewares"
	repository "studybuddy-backend-fast/api/repository/applicant"
	divrepo "studybuddy-backend-fast/api/repository/division"
	posrepo "studybuddy-backend-fast/api/repository/posisi"
	service "studybuddy-backend-fast/api/services/applicant"

	"github.com/gin-gonic/gin"
)

var (
	divisirepo divrepo.DivisiRepo = divrepo.NewDivisiRepo()
	posisrepo posrepo.PosisiRepo = posrepo.NewPosisiRepo()
	applicantRepo       repository.ApplicantRepo       = repository.NewApplicantRepo()
	applicantService    service.ApplicantService       = service.NewApplicantService(applicantRepo,divisirepo,posisrepo)
	applicantController controller.ApplicantController = controller.NewApplicantController(applicantService)
)

func findAllApplicant(c *gin.Context) {
	res := applicantController.FindAllApplicant(c)
	c.JSON(200, gin.H{"msg": "Query successful", "data": res})
}

func findApplicantByID(c *gin.Context) {
	res := applicantController.FindApplicantByID(c)
	c.JSON(200, gin.H{"msg": "Query successful", "data": res})
}

func createApplicant(c *gin.Context) {
	res := applicantController.CreateApplicant(c)
	c.JSON(200, gin.H{"msg": "Query successful", "data": res})
}

func createSolution(c *gin.Context) {
	res := applicantController.CreateSolution(c)
	c.JSON(200, gin.H{"msg": "Query successful", "data": res})
}

func updateApplicant(c *gin.Context) {
	res := applicantController.UpdateApplicant(c)
	c.JSON(200, gin.H{"msg": "Query successful", "data": res})
}

func ApplicantRoutes(router *gin.RouterGroup) {
	route := router.Group("/applicant")
	authMiddleware, _ := auth.New("applicant-access")
	{
		router.POST("/auth", authMiddleware.LoginHandler)
		route.GET("/", applicantController.ListApplicant)
		route.GET("/:id", findApplicantByID)
		route.PUT("/:id", updateApplicant)
		route.POST("/", createApplicant)
		// Authorized routes
		applicantRouter := router.Group("/")
		applicantRouter.Use(authMiddleware.MiddlewareFunc())
		{
			applicantRouter.POST("/applicant/question")
			applicantRouter.POST("/applicant/solution", createSolution)
		}
		applicant := router.Group("/applicants")
		{
			applicant.POST("/email",applicantController.Email)
			applicant.PUT("/update/:id",applicantController.Update)
		}
	}
}
