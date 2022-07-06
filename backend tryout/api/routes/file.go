package routes

import (
	"github.com/gin-gonic/gin"
	controller "studybuddy-backend-fast/api/controller/file"
)

var fileController controller.FileController = controller.NewFileController()

func FileRoutes(route *gin.RouterGroup) {
	router := route.Group("/file")
	{
		router.POST("/", fileController.HandleUpload)
		router.GET("/:file", fileController.SendFileUpload)
	}
}
