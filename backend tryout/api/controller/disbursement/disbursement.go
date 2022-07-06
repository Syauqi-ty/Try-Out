package controller

import (
	entity "studybuddy-backend-fast/api/entity"
	service "studybuddy-backend-fast/api/services/disbursement"

	"github.com/gin-gonic/gin"
)

type DisbursementController interface {
	FindAllDisbursement(c *gin.Context)
	CreateDisbursement(ctx *gin.Context) entity.Disbursement
}

type disbursementController struct {
	service service.DisbursementService
}

func NewDisbursementController(service service.DisbursementService) DisbursementController {
	return &disbursementController{
		service: service,
	}
}

func (s *disbursementController) CreateDisbursement(ctx *gin.Context) entity.Disbursement {
	newDisbursement := entity.Disbursement{}
	if err := ctx.ShouldBind(&newDisbursement); err != nil {
		return newDisbursement
	}
	return s.service.CreateDisbursement(newDisbursement)
}

func (s *disbursementController) FindAllDisbursement(ctx *gin.Context) {
	res := s.service.FindAllDisbursement()
	ctx.JSON(200, gin.H{"msg": "Query successful", "data": res})
}
