package routes

import (
	"api-blog-go/controller"

	"github.com/gin-gonic/gin"
)

// DefineProductRoutes configura as rotas de produtos
func DefineProductRoute(router *gin.Engine, productController *controller.ProductController) {
	v1 := router.Group("/api/")

	{
		v1.GET("/products", productController.GetProducts)
		v1.POST("/create-product", productController.CreateProduct)
		v1.GET("/product/:productId", productController.GetProductById) // get a product for id
	}
}
