package routers

import (
	controllers "final-project-4/httpserver/controllers"
	"final-project-4/httpserver/middleware"
	"final-project-4/utils"

	"github.com/gin-gonic/gin"
)

func TransactionHistoryRouter(route *gin.RouterGroup, transactionHistoryController controllers.TransactionHistoryController, authService utils.AuthHelper) *gin.RouterGroup {
	transactionRouter := route.Group("/transactions")
	{
		transactionRouter.POST("", middleware.JwtGuard(authService), transactionHistoryController.CreateTransaction)
		transactionRouter.GET("user-transactions", middleware.JwtGuard(authService),
			transactionHistoryController.GetUserTransactions)
		transactionRouter.GET("my-transactions", middleware.JwtGuard(authService), transactionHistoryController.GetMyTransactions)
	}
	return transactionRouter
}
