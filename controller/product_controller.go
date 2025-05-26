package controller

import (
	"api-blog-go/model"
	"api-blog-go/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	productUseCase usecase.ProductUsecase
}

func NewProductController(usecase usecase.ProductUsecase) ProductController {
	return ProductController{
		productUseCase: usecase,
	}
}

func (p *ProductController) GetProducts(ctx *gin.Context) {

	// validation of paramenters
	page := ctx.Query("page")
	limit := ctx.Query("limit")

	pageNumber, err := strconv.Atoi(page)

	if err != nil {
		responseMessage := model.Response{
			Message: "The 'page' parameter must be a number!",
		}

		ctx.JSON(http.StatusBadRequest, responseMessage)
	}

	limitNumber, err := strconv.Atoi(limit)

	if err != nil {
		responseMessage := model.Response{
			Message: "the 'limit' parameter must be a number",
		}

		ctx.JSON(http.StatusBadRequest, responseMessage)
	}

	products, count, err := p.productUseCase.GetProducts(pageNumber, limitNumber)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if products == nil {
		response := model.Response{
			Message: "No products found!",
		}
		ctx.JSON(http.StatusNoContent, response)
	}

	/* aqui eu tenho que declarar na forma de map, não basta abrir chaves!! */
	response := model.ProductResponse{
		Page:    page,
		PerPage: limit,
		Total:   count,
		Data:    products,
	}

	ctx.JSON(http.StatusOK, response)
}

func (p *ProductController) CreateProduct(ctx *gin.Context) {

	var listProducts []model.Products
	err := ctx.BindJSON(&listProducts)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	message, err := p.productUseCase.CreateProduct(listProducts)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, message)
}

// Get product by ID
func (p *ProductController) GetProductById(ctx *gin.Context) {
	id := ctx.Param("productId")
	if id == "" {
		response := model.Response{
			Message: "The id product must not null",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	product, err := p.productUseCase.GetProductById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if product == nil {
		response := model.Response{
			Message: "Product not found!",
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	ctx.JSON(http.StatusOK, product)
}
