package routes

import (
	"api-blog-go/controller"

	"github.com/gin-gonic/gin"
)

// DefineProductRoutes configura as rotas de produtos
func DefineTableRoute(router *gin.Engine, tableController *controller.TableController) {
	v1 := router.Group("/api")

	{
		v1.GET("/tables", tableController.GetTables)
		v1.POST("/create-table", tableController.CreateTable)
	}
}
