package route

import (
	controller "villa-akmali/api/controller/order"
	repository "villa-akmali/api/repository/order"
	repositoryyear "villa-akmali/api/repository/year"
	service "villa-akmali/api/service/order"

	"github.com/gin-gonic/gin"
)


var (
	yearrepo repositoryyear.YearRepo = repositoryyear.NewYearRepo()
	orderrepo repository.OrderRepo = repository.NewOrderRepo()
	orderservice service.OrderService = service.NewOrderService(orderrepo,userrepo,yearrepo,categoryrepo)
	ordercontroller controller.OrderController = controller.NewOrderService(orderservice)
)

func OrderRoutes(route *gin.RouterGroup) {
	router := route.Group("/order")
	{
		router.POST("/create",ordercontroller.CreateOrder)
		router.GET("/all",ordercontroller.FindAll)
		router.GET("/graphyear",ordercontroller.GraphYearly)
		router.GET("/graphmonth",ordercontroller.GraphMonth)
		router.GET("/graphweek",ordercontroller.GraphWeek)
		router.POST("/filter",ordercontroller.Filter)
		router.GET("/sumall",ordercontroller.SumAll)
		router.GET("/yearly",ordercontroller.Yearly)
		router.GET("/monthly",ordercontroller.Monthly)
		router.GET("/weekly",ordercontroller.Weekly)
		router.GET("/data/:id",ordercontroller.FindByID)
		router.PUT("/:id/update",ordercontroller.Update)
	}
}