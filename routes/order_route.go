package routes

import (
	"api-blog-go/controller"

	"github.com/gin-gonic/gin"
)

func DefineOrderRoute(router *gin.Engine, orderController *controller.OrderController) {
	v1 := router.Group("/api/orders")

	{
		v1.POST("/create-section", orderController.CreateSection)
		v1.POST("/create-order", orderController.CreateOrder)
	}
}
