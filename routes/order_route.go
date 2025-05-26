package routes

import (
	"api-blog-go/controller"

	"github.com/gin-gonic/gin"
)

func DefineOrderRoute(router *gin.Engine, orderController *controller.OrderController) {
	v1 := router.Group("/api/orders")
	/* TODO: fazer rota de fechar sessão e fazer rota para atualizar o status de pedidos */
	{
		v1.POST("/create-section", orderController.CreateSection)
		v1.POST("/create-order", orderController.CreateOrder)
		v1.GET("/get-orders/", orderController.GetOrders)
		v1.GET("/get-orders/:id_order", orderController.GetOrderById)
		v1.GET("/get-orders-itens", orderController.GetOrderItems)
		v1.GET("/get-orders-itens/:id_order", orderController.GetOrderItems)
	}
}
