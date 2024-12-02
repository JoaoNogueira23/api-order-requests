package controller

import (
	"api-blog-go/model"
	"api-blog-go/usecase"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
	orderUsecase usecase.OrderUsecase
}

func NewOrderController(usecase usecase.OrderUsecase) OrderController {
	return OrderController{
		orderUsecase: usecase,
	}
}

func (o *OrderController) CreateSection(ctx *gin.Context) {
	var payload model.OrderCreateSectionRq
	err := ctx.BindJSON(&payload)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	_, err = o.orderUsecase.CreateSection(payload.IdTable)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusCreated, payload.IdTable)

}

func (o *OrderController) CreateOrder(ctx *gin.Context) {
	var payload model.OrderCreateRq
	err := ctx.BindJSON(&payload)

	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	rowsEffected, err := o.orderUsecase.CreateOrder(payload.ProductsList, payload.IdSection)

	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusCreated, model.Response{
		Message: fmt.Sprintf("Pedido registrado com sucesso! Foram requisitados %d produtos", &rowsEffected),
	})

}

func (o *OrderController) GetOrders(ctx *gin.Context) {
	var payload model.GetOrdersRq
	err := ctx.BindJSON(&payload)

	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, err)
	}

}
