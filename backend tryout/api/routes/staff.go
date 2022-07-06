package routes

import (
	controller "studybuddy-backend-fast/api/controller/staff"
	auth "studybuddy-backend-fast/api/middlewares"
	repository "studybuddy-backend-fast/api/repository/staff"
	service "studybuddy-backend-fast/api/services/staff"

	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	staffrepo       repository.StaffRepo       = repository.NewStaffRepo()
	staffservice    service.StaffService       = service.NewStaffService(staffrepo)
	staffcontroller controller.StaffController = controller.NewStaffController(staffservice)
)

func staffone(ctx *gin.Context) {
	ctx.JSON(200, staffcontroller.FindOneById(ctx))
}

func staffaccess(ctx *gin.Context) {
	ctx.JSON(200, staffcontroller.FindByAccess(ctx))
}

func staffaccesscheck(ctx *gin.Context) {
	ctx.JSON(200, staffcontroller.FindByAccessCheck(ctx))
}

func editstaff(ctx *gin.Context) {
	err := staffcontroller.Update(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"messages": "Succesfully Edit Staff"})
	}
}

func createstaff(ctx *gin.Context) {
	err := staffcontroller.Save(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"messages": "Succesfully Create Staff"})
	}
}

func deletestaff(ctx *gin.Context) {
	err := staffcontroller.Delete(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"messages": "Succesfully Delete Staff"})
	}
}

func StaffRoutes(route *gin.RouterGroup) {
	router := route.Group("/staff")
	authMiddleware, _ := auth.New("staff-access")
	{
		router.GET("/access/:access", staffaccess)
		router.GET("/byid/:id", staffone)
		router.GET("/accesscheck/:id", staffaccesscheck)

		router.GET("/", staffcontroller.FindAll)
		router.PUT("/", editstaff)
		router.DELETE("/", deletestaff)

		adminRouter := router.Group("/")
		adminRouter.Use(authMiddleware.MiddlewareFunc())
		{
			adminRouter.POST("/", createstaff)
		}
	}
}
