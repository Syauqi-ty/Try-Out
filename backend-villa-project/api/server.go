package api

import (
	"villa-akmali/api/route"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)


func setupRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(gin.Logger())

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AddAllowHeaders("Access-Control-Allow-Origin")
	config.AddAllowHeaders("Authorization")

	router.Use(cors.New(config))
	grouping := router.Group("/api")
	{
		route.UserRoutes(grouping)
		route.CategoryRoutes(grouping)
		route.OrderRoutes(grouping)
	}
	return router
}

func Run() {
	router := setupRouter()
	router.Run(viper.GetString(`server.address`))
}