package controller

import (
	entity "studybuddy-backend-fast/api/entity"
	service "studybuddy-backend-fast/api/services/posisi"

	"github.com/gin-gonic/gin"
)
type PosisiController interface {
	CreatePosisi(c *gin.Context)
}

type posisiController struct {
	service service.PosisiService
}
func NewPosisiController(service service.PosisiService) PosisiController {
	return &posisiController{service: service}
}

func (s *posisiController) CreatePosisi(c *gin.Context) {
	var posisi entity.Posisi
	if err := c.ShouldBind(&posisi); err == nil {
		res := s.service.CreatePosisi(posisi)
		if res == 0 {
			c.JSON(400, gin.H{"msg": res})
		} else{
			c.JSON(200, gin.H{"msg": "Succesfully Posted", "data": res})
		}
	} else {
		c.JSON(400, gin.H{"msg": err.Error()})
	}
}