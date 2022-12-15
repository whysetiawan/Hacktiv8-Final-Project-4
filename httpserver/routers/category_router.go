package routers

import (
	controllers "final-project-4/httpserver/controllers"
	"final-project-4/httpserver/middleware"
	"final-project-4/utils"

	"github.com/gin-gonic/gin"
)

func CategoryRouter(route *gin.RouterGroup, categoryController controllers.CategoryController, authService utils.AuthHelper) *gin.RouterGroup {
	categoryRouter := route.Group("/categories")
	{
		categoryRouter.POST("", middleware.JwtGuard(authService), categoryController.CreateCategory)
		categoryRouter.GET("", middleware.JwtGuard(authService), categoryController.GetCategories)
		categoryRouter.PATCH(":categoryid", middleware.JwtGuard(authService), categoryController.UpdateCategory)
		categoryRouter.DELETE(":categoryid", middleware.JwtGuard(authService), categoryController.DeleteCategory)
	}
	return categoryRouter
}
