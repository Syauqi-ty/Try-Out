package controller

import (
	"strconv"
	entity "studybuddy-backend-fast/api/entity"
	service "studybuddy-backend-fast/api/services/applicant"

	"github.com/gin-gonic/gin"
)

type ApplicantController interface {
	FindAllApplicant(c *gin.Context) []entity.Applicant
	FindApplicantByID(c *gin.Context) entity.Applicant
	CreateApplicant(c *gin.Context) entity.Applicant
	CreateSolution(c *gin.Context) entity.Solution
	UpdateApplicant(c *gin.Context) entity.Applicant
	DeleteApplicant(c *gin.Context) error
	ListApplicant(c *gin.Context)
	Update(ctx *gin.Context)
	Email(ctx *gin.Context)
}

type applicantController struct {
	service service.ApplicantService
}

type deleteUserBody struct {
	id int
}

func NewApplicantController(service service.ApplicantService) ApplicantController {
	return &applicantController{service: service}
}

func (c *applicantController) FindAllApplicant(ctx *gin.Context) []entity.Applicant {
	return c.service.FindAllApplicant()
}

func (c *applicantController) FindApplicantByID(ctx *gin.Context) entity.Applicant {
	applicantID := ctx.Param("id")
	applicant, _ := strconv.Atoi(applicantID)
	return c.service.FindApplicantByID(applicant)

}

func (c *applicantController) CreateApplicant(ctx *gin.Context) entity.Applicant {
	newApplicant := entity.Applicant{}
	if err := ctx.ShouldBind(&newApplicant); err != nil {
		return newApplicant
	}
	return c.service.CreateApplicant(newApplicant)
}

func (c *applicantController) UpdateApplicant(ctx *gin.Context) entity.Applicant {
	newApplicantData := entity.Applicant{}
	if err := ctx.ShouldBind(&newApplicantData); err != nil {
		return newApplicantData
	}
	return c.service.UpdateApplicant(newApplicantData)
}

func (c *applicantController) DeleteApplicant(ctx *gin.Context) error {
	var applicant entity.Applicant
	err := ctx.ShouldBind(&applicant)
	c.service.DeleteApplicant(applicant)
	if err == nil {
		return nil
	} else {
		return err
	}
}

func (c *applicantController) CreateSolution(ctx *gin.Context) entity.Solution {
	newSolution := entity.Solution{}
	if err := ctx.ShouldBind(&newSolution); err != nil {
		return newSolution
	}
	return c.service.CreateSolution(newSolution)
}
func (c *applicantController) ListApplicant(ctx *gin.Context) {
	data := c.service.ListApplicantNew()
	ctx.JSON(200,gin.H{"data":data})
}

func (c *applicantController) Update(ctx *gin.Context){
	var applicant entity.Applicant
	id,_ := strconv.Atoi(ctx.Param("id"))
	if err := ctx.ShouldBind(&applicant); err != nil{
		c.service.Update(id,applicant)
		ctx.JSON(200,"Succesfully Updated")
	} else {
		ctx.JSON(400,"Yang Bener Aja Dong")
	}
}

func (c *applicantController) Email(ctx *gin.Context){
	var applicant entity.Applicant
	ctx.ShouldBindJSON(&applicant)
	c.service.Email(applicant)
	ctx.JSON(200,"Email Sent")
}