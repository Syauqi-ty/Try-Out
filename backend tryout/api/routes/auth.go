package routes

import (
	authController "studybuddy-backend-fast/api/controller/auth"
	auth "studybuddy-backend-fast/api/middlewares"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(route *gin.RouterGroup) {
	router := route.Group("/auth")
	authMiddleware, _ := auth.New("logged-in-access")
	{
		router.POST("/", authMiddleware.LoginHandler)
		router.GET("/", authMiddleware.RefreshHandler)
	}

	oauthRouter := route.Group("/oauth")
	{
		oauthRouter.GET("/", authController.GetAuthURL)
		oauthRouter.POST("/callback", authController.HandleGoogleCallback)
		oauthRouter.GET("/verify", authController.AuthFromAccessToken)
		oauthRouter.GET("/callback", authController.HandleGoogleCallback)
	}
}
