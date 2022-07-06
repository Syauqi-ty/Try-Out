package controller

import (
	authService "studybuddy-backend-fast/api/services/auth"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var timezone string = viper.GetString("timezone")

type GetOAuthURLRequest struct {
	Code  string `json:"code" form:"code"`
	State string `json:"state" form:"state"`
}

func GetAuthURL(c *gin.Context) {
	res := authService.GetAuthURL()
	c.Redirect(307, res)
}

func HandleGoogleCallback(c *gin.Context) {
	var req GetOAuthURLRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(400, gin.H{"msg": "Invalid request"})
	}
	student := authService.HandleGoogleCallback(req.State, req.Code)
	loc, _ := time.LoadLocation(timezone)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":      time.Now().In(loc).Add(time.Hour).Unix(),
		"orig_iat": time.Now().In(loc).Unix(),
		"status":   "student",
		"user_id":  student.ID,
	})
	tokenString, _ := token.SignedString([]byte(viper.GetString("secret")))

	if student.ID == 0 {
		c.Redirect(301, viper.GetString("server.name")+"/register")
	} else {
		c.HTML(200, "oauth.tmpl", gin.H{"name": student.Name, "token": tokenString, "user_id": student.ID})
	}
}

func AuthFromAccessToken(c *gin.Context) {
	accessToken := c.DefaultQuery("access_token", "")
	provider := c.DefaultQuery("provider", "google")

	if accessToken == "" {
		c.JSON(400, gin.H{"msg": "Bad request"})
	} else {
		if token, student := authService.VerifyAccessToken(accessToken, provider); token != "" {
			c.JSON(200, gin.H{"token": token, "data": student})
		} else {
			c.JSON(400, gin.H{"msg": "Bad request"})
		}
	}
}
