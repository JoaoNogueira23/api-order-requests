package controller

import (
	"api-blog-go/model"
	"api-blog-go/usecase"
	"api-blog-go/websocket"
	"encoding/json"
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

	// send websocket message
	response := map[string]any{
		"message": fmt.Sprintf("Pedido registrado com sucesso! Foram requisitados %d produtos na seção %s", rowsEffected, payload.IdSection),
		"data": map[string]any{
			"id_section":     payload.IdSection,
			"products_count": rowsEffected,
		},
	}
	// Marshal the response to JSON
	jsonBytes, err := json.Marshal(response)
	if err != nil {
		fmt.Println("Error on json message:", err)
	}

	websocket.Broadcast <- jsonBytes // Envia a mensagem para o canal de broadcast

	ctx.JSON(http.StatusCreated, model.Response{
		Message: fmt.Sprintf("Pedido registrado com sucesso! Foram requisitados %d produtos", &rowsEffected),
	})

}

func (o *OrderController) GetOrders(ctx *gin.Context) {
	id := ctx.Query("id_table")

	if id == "" {
		response := model.Response{
			Message: "The id of table do not null",
		}
		ctx.JSON(http.StatusBadRequest, response)
	}

	orders, err := o.orderUsecase.GetOrders(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	if orders == nil {
		response := model.Response{
			Message: "No orders found!",
		}
		ctx.JSON(http.StatusNoContent, response)
	}

	ctx.JSON(http.StatusOK, orders)

}

func (o *OrderController) GetOrderItems(ctx *gin.Context) {
	idOrder := ctx.Query("id_order")

	if idOrder == "" {
		response := model.Response{
			Message: "The id of order do not null",
		}
		ctx.JSON(http.StatusBadRequest, response)
	}

	orderItens, err := o.orderUsecase.GetOrderItems(idOrder)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	if orderItens == nil {
		response := model.Response{
			Message: "No order items found!",
		}
		ctx.JSON(http.StatusNoContent, response)
	}

	ctx.JSON(http.StatusOK, orderItens)
}

func (o *OrderController) GetOrderById(ctx *gin.Context) {
	idOrder := ctx.Query("id_order")

	if idOrder == "" {
		response := model.Response{
			Message: "The id of order do not null",
		}
		ctx.JSON(http.StatusBadRequest, response)
	}

	order, err := o.orderUsecase.GetOrderById(idOrder)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	if order == nil {
		response := model.Response{
			Message: "No order found!",
		}
		ctx.JSON(http.StatusNotFound, response)
	}

	ctx.JSON(http.StatusOK, order)
}
