package controller

import (
	"strconv"
	entity "studybuddy-backend-fast/api/entity"
	mapper "studybuddy-backend-fast/api/helper/querystring"
	sservice "studybuddy-backend-fast/api/services/score"
	service "studybuddy-backend-fast/api/services/test"

	jwt "github.com/appleboy/gin-jwt/v2"

	"github.com/gin-gonic/gin"
)

type TestController interface {
	FindAllTest(ctx *gin.Context)
	FindTest(ctx *gin.Context) entity.TestQuestionMin
	CreateTest(ctx *gin.Context) entity.Test
	UpdateTest(ctx *gin.Context) entity.Test
	DeleteTest(ctx *gin.Context) error

	FindAllTestMin(ctx *gin.Context) []entity.TestMin
	FindTestMin(ctx *gin.Context) entity.TestMin
	FindTestFull(ctx *gin.Context)

	AnswerQuestion(ctx *gin.Context)
	LastBattle(ctx *gin.Context)
	IkutTest(ctx *gin.Context)
	AvailableTest(ctx *gin.Context)
	SoalBattle(ctx *gin.Context)
}

type testController struct {
	service      service.TestService
	scoreService sservice.ScoreService
}

type AnswerBody struct {
	Answer string `json:"answer;required"`
}

func NewTestController(service service.TestService, sservice sservice.ScoreService) TestController {
	return &testController{service, sservice}
}

func (c *testController) AnswerQuestion(ctx *gin.Context) {
	var answer AnswerBody
	claims := jwt.ExtractClaims(ctx)
	testID, _ := strconv.Atoi(ctx.Param("testID"))
	studentID := int(claims["user_id"].(float64))
	questionID, _ := strconv.Atoi(ctx.Param("questionID"))

	if err := ctx.ShouldBind(&answer); err == nil {
		c.scoreService.AnswerTestQuestion(testID, questionID, studentID, answer.Answer)
		ctx.JSON(200, gin.H{"msg": "Answer saved"})
	} else {
		ctx.JSON(400, gin.H{"msg": "Bad request, missing fields"})
	}
}

func (c *testController) FindAllTest(ctx *gin.Context) {
	pagination := mapper.GetTestMapper(ctx)
	res := c.service.FindAllTest(pagination)
	pnum := strconv.Itoa(pagination["page"])
	lnum := strconv.Itoa(pagination["limit"])
	ctx.JSON(200, gin.H{"msg": "On page " + pnum + " with limit " + lnum, "data": res})
}

func (c *testController) FindTest(ctx *gin.Context) entity.TestQuestionMin {
	param := ctx.Param("id")
	return c.service.FindTest(param)
}

func (c *testController) FindTestFull(ctx *gin.Context) {
	res := c.service.FindTestFull(ctx.Param("id"))
	ctx.JSON(200, gin.H{"msg": "Query successful", "data": res})
}

func (c *testController) CreateTest(ctx *gin.Context) entity.Test {
	var test entity.Test
	err := ctx.ShouldBind(&test)
	if err != nil {
		return test
	}
	return c.service.CreateTest(test)
}

func (c *testController) UpdateTest(ctx *gin.Context) entity.Test {
	newTestData := entity.Test{}
	if err := ctx.ShouldBind(&newTestData); err != nil {
		return newTestData
	}
	c.service.UpdateTest(newTestData)
	return newTestData
}

func (c *testController) DeleteTest(ctx *gin.Context) error {
	var test entity.Test
	err := ctx.ShouldBind(&test)
	c.service.DeleteTest(test)
	if err != nil {
		return nil
	}
	return err
}

func (c *testController) FindAllTestMin(ctx *gin.Context) []entity.TestMin {
	return c.service.FindAllTestMin()
}

func (c *testController) FindTestMin(ctx *gin.Context) entity.TestMin {
	param := ctx.Param("id")
	return c.service.FindTestMin(param)
}

func (c *testController) AvailableTest(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	res := c.service.Available(id)
	ctx.JSON(200, gin.H{"data": res})
}

func (c *testController) LastBattle(ctx *gin.Context) {
	var test entity.Test
	res := c.service.LastBattle(test)
	ctx.JSON(200, gin.H{"data": res})
}
func (c *testController) IkutTest(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	res := c.service.BattleIkut(id)
	ctx.JSON(200, gin.H{"data": res})
}

func (c *testController) SoalBattle(ctx *gin.Context) {
	id,_ := strconv.Atoi(ctx.Param("id"))
	subtest := ctx.Param("subtest")
	data := c.service.SoalBattle(id,subtest)
	ctx.JSON(200,gin.H{"data":data})
}
