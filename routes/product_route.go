package routes

import (
	"api-blog-go/controller"

	"github.com/gin-gonic/gin"
)

// DefineProductRoutes configura as rotas de produtos
func DefineProductRoute(router *gin.Engine, productController *controller.ProductController) {
	router.GET("/products", productController.GetProducts)
	router.POST("/product", productController.CreateProduct)
	router.GET("/product/:productId", productController.GetProductById)
}
