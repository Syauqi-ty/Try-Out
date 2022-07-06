package routes

import (
	controller "studybuddy-backend-fast/api/controller/payment"
	auth "studybuddy-backend-fast/api/middlewares"
	repository "studybuddy-backend-fast/api/repository/payment"
	service "studybuddy-backend-fast/api/services/payment"

	"github.com/gin-gonic/gin"
)

var (
	paymentRepo       repository.PaymentRepo       = repository.NewPaymentRepo()
	paymentService    service.PaymentService       = service.NewPaymentService(paymentRepo, testRepo)
	paymentController controller.PaymentController = controller.NewPaymentController(paymentService)
)

func PaymentRoutes(route *gin.RouterGroup) {
	router := route.Group("/payment")
	authMiddleware, _ := auth.New("student-access")
	// router.GET("/", paymentController.FindAllPayment)

	// EWALLET RELATED
	router.POST("/ovo/callback", paymentController.OVOCallback)
	router.POST("/dana/callback", paymentController.DanaCallback)
	router.POST("/linkaja/callback", paymentController.LinkAjaCallback)

	// Authorized Routes
	studentRouter := router.Group("/")
	studentRouter.Use(authMiddleware.MiddlewareFunc())
	{
		studentRouter.POST("/ewallet", paymentController.RequestEwalletPayment)
	}

}
