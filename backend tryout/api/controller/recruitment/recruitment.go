package controller

import (
	"strconv"
	"github.com/gin-gonic/gin"
	entity "studybuddy-backend-fast/api/entity"
	service "studybuddy-backend-fast/api/services/recruitment"
)

type RecruitmentController interface {
	CreateSoal(c *gin.Context)
	FindSoalByID(c *gin.Context)
}

type recruitmentController struct {
	service service.RecruitmentService
}
func NewRecruitmentController(s service.RecruitmentService) RecruitmentController {
	return &recruitmentController{s}
}

type Body struct {
	Pendaftar entity.Applicant `json:"pendaftar"`
	Soal []entity.Recruitment `json:"soal"`
}
func (s *recruitmentController) CreateSoal(c *gin.Context) {
	var body Body
	bodyjson := c.ShouldBind(&body)
	if bodyjson == nil {
		data := s.service.CreateSoal(body.Pendaftar,body.Soal)
		c.JSON(200,gin.H{"status":data})
	} else{
		c.JSON(400, gin.H{"msg": bodyjson.Error()})
	}
}
func (s *recruitmentController) FindSoalByID(c *gin.Context){
	id,_ := strconv.Atoi(c.Param("id"))
	data := s.service.FindSoalByApplicantID(id)
	c.JSON(200,gin.H{"data":data})
}