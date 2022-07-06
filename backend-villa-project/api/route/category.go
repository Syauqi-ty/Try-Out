package route

import (
	controller "villa-akmali/api/controller/category"
	repository "villa-akmali/api/repository/category"
	service "villa-akmali/api/service/category"

	"github.com/gin-gonic/gin"
)


var (
	categoryrepo repository.CategoryRepo = repository.NewCategoryRepo()
	categoryservice service.CategoryService = service.NewCategoryService(categoryrepo)
	categorycontroller controller.CategoryController = controller.NewCategoryController(categoryservice)
)

func CategoryRoutes(route *gin.RouterGroup) {
	router := route.Group("/category")
	{
		router.POST("/create",categorycontroller.CreateCategory)
		router.GET("/all",categorycontroller.FindAll)
	}
}