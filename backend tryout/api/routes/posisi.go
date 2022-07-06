package routes

import (
	controller "studybuddy-backend-fast/api/controller/posisi"
	repository "studybuddy-backend-fast/api/repository/posisi"
	service "studybuddy-backend-fast/api/services/posisi"

	"github.com/gin-gonic/gin"
)

var (
	posisirepo repository.PosisiRepo = repository.NewPosisiRepo()
	posisiservice service.PosisiService = service.NewPosisiService(posisirepo)
	posisicontroller controller.PosisiController = controller.NewPosisiController(posisiservice)
)

func PosisiRoutes(route *gin.RouterGroup){
	router := route.Group("/posisi")
	{
		router.POST("/create",posisicontroller.CreatePosisi)
	}
}