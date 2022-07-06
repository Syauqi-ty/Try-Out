package routes

import (
	controller "studybuddy-backend-fast/api/controller/division"
	repository "studybuddy-backend-fast/api/repository/division"
	posrepo "studybuddy-backend-fast/api/repository/posisi"
	service "studybuddy-backend-fast/api/services/division"

	"github.com/gin-gonic/gin"
)

var (
	divisionrepo repository.DivisiRepo = repository.NewDivisiRepo()
	positionrepo posrepo.PosisiRepo = posrepo.NewPosisiRepo()
	divisionservice service.DivisionService = service.NewDivisionService(divisionrepo,positionrepo)
	divisioncontroller controller.DivisionController = controller.NewDivisionController(divisionservice)
)

func DivisionRoutes(route *gin.RouterGroup){
	router := route.Group("/division")
	{
		router.GET("/",divisioncontroller.FindAll)
		router.POST("/create",divisioncontroller.CreateDivision)
		router.GET("/:divid/desc/:posid",divisioncontroller.Description)
	}
}