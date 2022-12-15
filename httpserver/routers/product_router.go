package routers

import (
	controllers "final-project-4/httpserver/controllers"
	"final-project-4/httpserver/middleware"
	"final-project-4/utils"

	"github.com/gin-gonic/gin"
)

func ProductRouter(route *gin.RouterGroup, productController controllers.ProductController, authService utils.AuthHelper) *gin.RouterGroup {
	productRouter := route.Group("/product")
	{
		productRouter.POST("create", productController.CreateProduct)
		productRouter.GET("", middleware.JwtGuard(authService), productController.GetProducts)
		productRouter.PUT(":productid", middleware.JwtGuard(authService), productController.UpdateProduct)
		productRouter.DELETE(":productid", middleware.JwtGuard(authService), productController.DeleteProduct)
	}
	return productRouter
}
