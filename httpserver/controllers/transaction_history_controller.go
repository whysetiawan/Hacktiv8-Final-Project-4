package controllers

import (
	"errors"
	"final-project-4/httpserver/dto"
	"final-project-4/httpserver/models"
	"final-project-4/httpserver/services"
	"final-project-4/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TransactionHistoryController interface {
	CreateTransaction(ctx *gin.Context)
	GetUserTransactions(ctx *gin.Context)
	GetMyTransactions(ctx *gin.Context)
}

type transactionHistoryController struct {
	transactionService services.TransactionHistoryService
}

func NewTransactionHistoryController(transactionService services.TransactionHistoryService) *transactionHistoryController {
	return &transactionHistoryController{
		transactionService,
	}
}

// CreateTransaction godoc
//	@Tags		Transactions
//	@Summary	create a transaction
//	@Param		Transaction	body		dto.CreateTransactionDto	true	"CreateTransactionDto"
//	@Success	201			{object}	utils.HttpSuccess[models.TransactionHistoryModel]
//	@Failure	400			{object}	utils.HttpError
//	@Failure	500			{object}	utils.HttpError
//	@Router		/transactions [post]
//	@Security	BearerAuth
func (c *transactionHistoryController) CreateTransaction(ctx *gin.Context) {

	var dto dto.CreateTransactionDto
	err := ctx.BindJSON(&dto)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewHttpError("Bad Request", err.Error()))
		return
	}

	userCredential, isExist := ctx.Get("user")

	if !isExist {
		ctx.JSON(http.StatusBadRequest, utils.NewHttpError("Bad Request", errors.New("invalid credential")))
		return
	}

	userModel := userCredential.(models.UserModel)

	transaction, err := c.transactionService.CreateTransaction(&dto, &userModel)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.NewHttpError("Internal Server Error", err.Error()))
		return
	}

	ctx.JSON(http.StatusCreated, utils.NewHttpSuccess("Category Created", transaction))
}

// GetUserTransactions godoc
//	@Tags		Transactions
//	@Summary	get all current user Transactions
//	@Success	200	{object}	utils.HttpSuccess[[]models.TransactionHistoryModel]
//	@Failure	400	{object}	utils.HttpError
//	@Failure	500	{object}	utils.HttpError
//	@Router		/transactions/my-transactions [get]
//	@Security	BearerAuth
func (c *transactionHistoryController) GetMyTransactions(ctx *gin.Context) {

	userCredential, isExist := ctx.Get("user")

	if !isExist {
		ctx.JSON(http.StatusBadRequest, utils.NewHttpError("Bad Request", errors.New("invalid credential")))
		return
	}

	userModel := userCredential.(models.UserModel)

	transactions, err := c.transactionService.GetMyTransactions(&userModel)

	// category, err := c.transactionService.GetUserTransaction(&userModel)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.NewHttpError("Internal Server Error", err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, utils.NewHttpSuccess("Success Get Transactions", transactions))
}

// GetUserTransactions godoc
//	@Tags		Transactions
//	@Summary	get all Transactions for admin
//	@Success	200	{object}	utils.HttpSuccess[[]models.TransactionHistoryModel]
//	@Failure	400	{object}	utils.HttpError
//	@Failure	500	{object}	utils.HttpError
//	@Router		/transactions/user-transactions [get]
//	@Security	BearerAuth
func (c *transactionHistoryController) GetUserTransactions(ctx *gin.Context) {

	userCredential, isExist := ctx.Get("user")

	if !isExist {
		ctx.JSON(http.StatusBadRequest, utils.NewHttpError("Bad Request", errors.New("invalid credential")))
		return
	}

	userModel := userCredential.(models.UserModel)

	if userModel.Role != "admin" {
		ctx.JSON(http.StatusForbidden, utils.NewHttpError("Access Forbidden", "you are not an admin"))
		return
	}

	transactions, err := c.transactionService.GetUserTransactions(&userModel)

	// category, err := c.transactionService.GetUserTransaction(&userModel)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.NewHttpError("Internal Server Error", err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, utils.NewHttpSuccess("Success Get Transactions", transactions))
}
