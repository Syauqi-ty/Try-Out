package controller

import (
	"strconv"
	entity "studybuddy-backend-fast/api/entity"
	service "studybuddy-backend-fast/api/services/score"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

type ScoreController interface {
	FindAllScore(c *gin.Context) []entity.ScoreMin
	UserStats(c *gin.Context)
	TestStats(c *gin.Context)
	AllStudentScore(c *gin.Context)
	CheckTest(c *gin.Context)
	GetRank(c *gin.Context)
	LeaderBoard(c *gin.Context)
	Graphic(c *gin.Context)
	AvgTKA(c *gin.Context)
	Battled(c *gin.Context)
	HighScore(c *gin.Context)
	HighScoreBattle(c *gin.Context)
	GraphicData(c *gin.Context)
	Quartile(c *gin.Context)
	Distribution(c *gin.Context)
	NewLeaderBoard(c *gin.Context)
	Rara(c *gin.Context)
	BenarSalah(c *gin.Context)
	Jawab(ctx *gin.Context)
}

type scoreController struct {
	service service.ScoreService
}

func NewScoreController(s service.ScoreService) ScoreController {
	return &scoreController{s}
}

func (s *scoreController) FindAllScore(c *gin.Context) []entity.ScoreMin {
	return s.service.FindAllScore()
}
func (s *scoreController) Distribution(c *gin.Context) {
	testid, _ := strconv.Atoi(c.Param("id"))
	data := s.service.Distribution(testid)
	c.JSON(200, gin.H{"data": data})
}
func (s *scoreController) BenarSalah(c *gin.Context){
	testid, _ := strconv.Atoi(c.Param("testid"))
	id,_ := strconv.Atoi(c.Param("id"))
	data := s.service.BenarSalah(testid,id)
	c.JSON(200, gin.H{"data": data})
}
func (s *scoreController) Rara(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data := s.service.Rara(id)
	c.JSON(200, gin.H{"data": data})
}
func (s *scoreController) Quartile(c *gin.Context) {
	testid, _ := strconv.Atoi(c.Param("id"))
	data := s.service.AllQuartil(testid)
	c.JSON(200, gin.H{"data": data})
}
func (s *scoreController) HighScore(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data := s.service.HighScore(id)
	c.JSON(200, gin.H{"data": data})
}
func (s *scoreController) GraphicData(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data := s.service.GraphicData(id)
	c.JSON(200, gin.H{"data": data})
}

func (s *scoreController) NewLeaderBoard(c *gin.Context) {
	testid, _ := strconv.Atoi(c.Param("id"))
	id, _ := strconv.Atoi(c.Param("studentID"))
	data := s.service.NewLeaderBoard(testid, id)
	c.JSON(200, gin.H{"data": data})
}

func (s *scoreController) GetRank(c *gin.Context) {
	testID, err1 := strconv.Atoi(c.Param("id"))
	studentID, err2 := strconv.Atoi(c.Param("studentID"))

	if err1 != nil || err2 != nil {
		c.JSON(400, gin.H{"msg": "Bad request, invalid IDs"})
	} else {
		res := s.service.GetRank(testID, studentID)
		c.JSON(200, gin.H{"msg": "Rank calculated", "data": res})
	}
}

func (s *scoreController) CheckTest(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	if testID, err := strconv.Atoi(c.Param("id")); err == nil {
		res := s.service.CheckAnswers(testID, int(claims["user_id"].(float64)))
		c.JSON(200, gin.H{"msg": "Check successful", "data": res})
	} else {
		c.JSON(400, gin.H{"msg": "Bad request, invalid test ID"})
	}
}
func (s *scoreController) AvgTKA(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	res := s.service.OutputAllAverage(id)
	c.JSON(200, gin.H{"data": res})
}
func (s *scoreController) AllStudentScore(c *gin.Context) {
	if id, err := strconv.Atoi(c.Param("id")); err == nil {
		if q := c.Query("test"); q == "" {
			res := s.service.StudentFScore(id)
			c.JSON(200, gin.H{"msg": "Scores fetched", "data": res})
		} else {
			t, _ := strconv.Atoi(q)
			res := s.service.StudentFScoreOfTest(id, t)
			c.JSON(200, gin.H{"msg": "Scores fetched", "data": res})
		}
	} else {
		c.JSON(400, gin.H{"msg": "Bad request, read the docs for references"})
	}
}
func (s *scoreController) HighScoreBattle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	res := s.service.HighScoreBattle(id)
	c.JSON(200, gin.H{"data": res})
}
func (s *scoreController) UserStats(c *gin.Context) {
	if id, err := strconv.Atoi(c.Param("id")); err == nil {
		res := s.service.CalculateStudentStats(id)
		c.JSON(200, gin.H{"msg": "Stats calculated", "data": res})
	} else {
		c.JSON(400, gin.H{"msg": "Bad request, read the docs for references"})
	}
}
func (s *scoreController) Battled(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	res := s.service.Battled(id)
	c.JSON(200, res)
}
func (s *scoreController) TestStats(c *gin.Context) {
	if id, err := strconv.Atoi(c.Param("id")); err == nil {
		res := s.service.CalculateTestStats(id)
		c.JSON(200, gin.H{"msg": "Stats calculated", "data": res})
	} else {
		c.JSON(400, gin.H{"msg": "Bad request, read the docs for references"})
	}
}

func (s *scoreController) LeaderBoard(c *gin.Context) {
	if test_id, err := strconv.Atoi(c.Param("id")); err == nil {
		if subtest := c.DefaultQuery("subtest", ""); subtest != "" {
			res := s.service.TypeLeaderBoard(test_id, subtest)
			c.JSON(200, gin.H{"msg": "Here you go", "data": res})
		} else {
			res := s.service.LeaderBoard(test_id)
			c.JSON(200, gin.H{"msg": "Here you go", "data": res})
		}
	} else {
		c.JSON(400, gin.H{"msg": "Bad request"})
	}
}
func (s *scoreController) Graphic(c *gin.Context) {
	var test entity.Test
	var score entity.Score
	pantat := s.service.Graphic(test, score)
	c.JSON(200, pantat)
}

func (s *scoreController) Jawab(ctx *gin.Context) {
	var req entity.Jawaban
	claims := jwt.ExtractClaims(ctx)
	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(400,gin.H{"msg":"Yang bener dong ah"})
	} else if req.UserID != int(claims["user_id"].(float64)){
		ctx.JSON(400,gin.H{"msg":"Saha Sia"})
	} else{
		ctx.JSON(200,gin.H{"data":s.service.Jawab(req)})
	}
}
