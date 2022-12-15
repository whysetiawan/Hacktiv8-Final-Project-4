package controllers

import (
	"errors"
	"final-project-4/httpserver/dto"
	"final-project-4/httpserver/models"
	"final-project-4/httpserver/services"
	"final-project-4/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductController interface {
	CreateProduct(ctx *gin.Context)
	GetProducts(ctx *gin.Context)
	UpdateProduct(ctx *gin.Context)
	DeleteProduct(ctx *gin.Context)
}

type productController struct {
	productService services.ProductService
	authService    utils.AuthHelper
}

func NewProductController(
	productService services.ProductService,
	authService utils.AuthHelper,
) *productController {
	return &productController{productService, authService}
}

// CreateProduct godoc
//	@Tags		Product
//	@Summary	create a product
//	@Param		Product	body		dto.InputProduct	true	"Create Category DTO"
//	@Success	201			{object}	utils.HttpSuccess[any]
//	@Failure	400			{object}	utils.HttpError
//	@Failure	500			{object}	utils.HttpError
//	@Router		/product/create [post]
//	@Security	BearerAuth
func (c *productController) CreateProduct(ctx *gin.Context) {
	var dto dto.InputProduct
	err := ctx.BindJSON(&dto)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewHttpError("Bad Request", err.Error()))
		return
	}

	data, _, err := c.productService.Create(dto)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.NewHttpError("Internal Server Error", err.Error()))
		return
	}

	ctx.JSON(http.StatusCreated, utils.NewHttpSuccess("product created", &data))
}

// GetProducts godoc
//	@Tags		Product
//	@Summary	get all product
//	@Success	200	{object}	utils.HttpSuccess[[]dto.ResponseProducts]
//	@Failure	400	{object}	utils.HttpError
//	@Failure	500	{object}	utils.HttpError
//	@Router		/product [get]
//	@Security	BearerAuth
func (c *productController) GetProducts(ctx *gin.Context) {
	userCredential, isExist := ctx.Get("user")
	_ = userCredential.(models.UserModel)

	if !isExist {
		ctx.JSON(http.StatusBadRequest, utils.NewHttpError("Bad Request", errors.New("invalid credential")))
		return
	}
	data, _, err := c.productService.GetProducts()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.NewHttpError("Internal Server Error", err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, utils.NewHttpSuccess("Get All Success", data))
}

// UpdateProduct godoc
//	@Tags		Product
//	@Summary	Update a product
//	@Param		id			path		int						true	"Product ID"
//	@Param		Product		body		dto.Product	true	"Update Category On User"
//	@Success	200			{object}	utils.HttpSuccess[dto.Product]
//	@Failure	400			{object}	utils.HttpError
//	@Failure	500			{object}	utils.HttpError
//	@Router		/product/{id} [put]
//	@Security	BearerAuth
func (c *productController) UpdateProduct(ctx *gin.Context) {
	var dto dto.Product
	err := ctx.BindJSON(&dto)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewHttpError("Bad Request", err.Error()))
		return
	}

	idParam := ctx.Param("productid")
	productID, err := strconv.Atoi(idParam)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewHttpError("Bad Request", err.Error()))
		return
	}

	userCredential, isExist := ctx.Get("user")

	if !isExist {
		ctx.JSON(http.StatusBadRequest, utils.NewHttpError("Bad Request", errors.New("invalid credential")))
		return
	}

	_ = userCredential.(models.UserModel)
	dto.ID = int64(productID)

	category, err := c.productService.UpdateProduct(&dto)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.NewHttpError("Internal Server Error", err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, utils.NewHttpSuccess("Update product Success", category))
}

// DeleteProduct godoc
//	@Tags		Product
//	@Summary	Delete a Product
//	@Param		id			path		int						true	"Product ID"
//	@Success  200 {object} utils.HttpSuccess[any]
//	@Failure	400			{object}	utils.HttpError
//	@Failure	500			{object}	utils.HttpError
//	@Router		/product/{id} [delete]
//	@Security	BearerAuth
func (c *productController) DeleteProduct(ctx *gin.Context) {

	idParam := ctx.Param("productid")
	productID, err := strconv.Atoi(idParam)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewHttpError("Bad Request", err.Error()))
		return
	}

	userCredential, isExist := ctx.Get("user")

	if !isExist {
		ctx.JSON(http.StatusBadRequest, utils.NewHttpError("Bad Request", errors.New("invalid credential")))
		return
	}
	var dto dto.Product

	_ = userCredential.(models.UserModel)
	dto.ID = int64(productID)

	_, err = c.productService.DeleteProduct(&dto)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.NewHttpError("Internal Server Error", err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, utils.NewHttpSuccess("product has been successfully deleted", ""))
}
