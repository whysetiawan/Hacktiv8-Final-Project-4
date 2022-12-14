package routers

import (
	controllers "final-project-4/httpserver/controllers"
	"final-project-4/httpserver/middleware"
	"final-project-4/utils"

	"github.com/gin-gonic/gin"
)

func UserRouter(route *gin.RouterGroup, userController controllers.UserController, authService utils.AuthHelper) *gin.RouterGroup {
	userRouter := route.Group("/user")
	{
		userRouter.POST("register", userController.Register)
		userRouter.POST("login", userController.Login)
		userRouter.PUT("", middleware.JwtGuard(authService), userController.UpdateUser)
		userRouter.DELETE("", middleware.JwtGuard(authService), userController.DeleteUser)
		userRouter.GET("", middleware.JwtGuard(authService), userController.GetUsers)
		userRouter.PATCH("topup", middleware.JwtGuard(authService), userController.TopUpBalance)
	}
	return userRouter
}
