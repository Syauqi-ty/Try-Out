package routes

import (
	controller "studybuddy-backend-fast/api/controller/disbursement"
	repository "studybuddy-backend-fast/api/repository/disbursement"
	service "studybuddy-backend-fast/api/services/disbursement"

	"github.com/gin-gonic/gin"
)

var (
	disbursementRepo       repository.DisbursementRepo       = repository.NewDisbursementRepo()
	disbursementService    service.DisbursementService       = service.NewDisbursementService(disbursementRepo)
	disbursementController controller.DisbursementController = controller.NewDisbursementController(disbursementService)
)

func CreateDisbursement(c *gin.Context) {
	res := disbursementController.CreateDisbursement(c)
	c.JSON(200, gin.H{"msg": "Query successful", "data": res})
}

func DisbursementRoutes(route *gin.RouterGroup) {
	router := route.Group("/disbursement")

	router.GET("/", disbursementController.FindAllDisbursement)
	router.POST("/", CreateDisbursement)
}
