package route

import (
	controller "villa-akmali/api/controller/user"
	repository "villa-akmali/api/repository/user"
	service "villa-akmali/api/service/user"

	"github.com/gin-gonic/gin"
)


var (
	userrepo repository.UserRepo = repository.NewUserRepo()
	userservice service.UserService = service.NewUserService(userrepo)
	usercontroller controller.UserController = controller.NewUserController(userservice)
)

func UserRoutes(route *gin.RouterGroup)  {
	router := route.Group("/user")
	{
		router.POST("/register",usercontroller.CreateUser)
		router.POST("/login",usercontroller.Login)
	}
}