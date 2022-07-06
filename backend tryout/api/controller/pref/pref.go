package controller

import (
	"github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"strconv"
	entity "studybuddy-backend-fast/api/entity"
	service "studybuddy-backend-fast/api/services/pref"
	testAuthService "studybuddy-backend-fast/api/services/testauth"
)

type PrefController interface {
	FindAllUni(c *gin.Context)
	FindAllUniMin(c *gin.Context)
	FindUniByID(c *gin.Context)
	FindUniMinByID(c *gin.Context)

	SetPref(c *gin.Context)
}

type prefController struct {
	serve           service.PrefService
	testAuthService testAuthService.TestAuthService
}

func NewPrefController(s service.PrefService, testAuthService testAuthService.TestAuthService) PrefController {
	return &prefController{s, testAuthService}
}

func (p *prefController) FindAllUni(c *gin.Context) {
	res := p.serve.FindAllFUni()
	c.JSON(200, gin.H{"msg": "Query successful", "data": res})
}

func (p *prefController) FindUniByID(c *gin.Context) {
	if id, err := strconv.Atoi(c.Param("id")); err == nil {
		if pid, err := strconv.Atoi(c.Query("prodi")); err == nil {
			res := p.serve.FindProdiOfUni(pid, id)
			c.JSON(200, gin.H{"msg": "Query successful", "data": res})
		} else {
			res := p.serve.FindFUniByID(id)
			c.JSON(200, gin.H{"msg": "Query successful", "data": res})
		}
	} else {
		c.JSON(400, gin.H{"msg": "Bad request, nothing fetched"})
	}
}

func (p *prefController) FindAllUniMin(c *gin.Context) {
	res := p.serve.FindAllUniMin()
	c.JSON(200, gin.H{"msg": "Query successful", "data": res})
}

func (p *prefController) FindUniMinByID(c *gin.Context) {
	if id, err := strconv.Atoi(c.Param("id")); err == nil {
		res := p.serve.FindUniMinByID(id)
		c.JSON(200, gin.H{"msg": "Query successful", "data": res})
	} else {
		c.JSON(400, gin.H{"msg": "Bad request, nothing fetched"})
	}
}

func (p *prefController) SetPref(c *gin.Context) {
	var setPrefBody entity.SetPrefBody
	claims := jwt.ExtractClaims(c)
	if err := c.ShouldBind(&setPrefBody); err != nil {
		c.JSON(400, gin.H{"msg": "Bad request"})
	} else {
		setPrefBody.StudentID = int(claims["user_id"].(float64))
		p.testAuthService.ChangePref(setPrefBody)
		c.JSON(200, gin.H{"msg": "Preferences saved"})
	}
}
