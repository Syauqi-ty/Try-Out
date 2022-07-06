package routes

import (
	controller "studybuddy-backend-fast/api/controller/recruitment"
	repository "studybuddy-backend-fast/api/repository/recruitment"
	apprepo "studybuddy-backend-fast/api/repository/applicant"
	service "studybuddy-backend-fast/api/services/recruitment"

	"github.com/gin-gonic/gin"
)

var (
	recruitmentrepo repository.RecruitmentRepo = repository.NewRecruitmentRepo()
	applicantrepo apprepo.ApplicantRepo = apprepo.NewApplicantRepo()
	recruitmentservice service.RecruitmentService = service.NewRecruitmentService(recruitmentrepo,applicantrepo)
	recruitmentcontroller controller.RecruitmentController = controller.NewRecruitmentController(recruitmentservice)
)

func RecruitmentRoutes(route *gin.RouterGroup){
	router := route.Group("/recruitment")
	{
		router.POST("/create",recruitmentcontroller.CreateSoal)
	}
	soal := route.Group("/soal")
	{
		soal.GET("/:id",recruitmentcontroller.FindSoalByID)
	}
}