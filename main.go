package main

import (
	"api-blog-go/controller"
	"api-blog-go/db"
	"api-blog-go/repository"
	"api-blog-go/routes"
	"api-blog-go/usecase"
	websocketapi "api-blog-go/websocket"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	// Configuração de CORS validation
	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	// ##### PRODUCT #####
	ProductRepository := repository.NewProductRepository(dbConnection) // database comunicate
	//Camada usecase
	ProductUseCase := usecase.NewProductUseCase(ProductRepository) // business logic
	//Camada de controllers
	ProductController := controller.NewProductController(ProductUseCase) // api controller
	// incluido a rota de produtos
	routes.DefineProductRoute(server, &ProductController)

	// ##### TABLES ######
	TableRepository := repository.NewTableRepository(dbConnection)
	TableUsecase := usecase.NewTableUsecase(TableRepository)
	TableController := controller.NewTableController(TableUsecase)

	// incluindo a rota de tables
	routes.DefineTableRoute(server, &TableController)

	// ##### ORDERS #####
	OrdersRepository := repository.NewOrderRepository(dbConnection)
	OrdersUsecase := usecase.NewOrderUsecase(OrdersRepository)
	OrdersController := controller.NewOrderController(OrdersUsecase)

	// incluindo a rota de orders
	routes.DefineOrderRoute(server, &OrdersController)

	// ##### WEBSOCKET #####
	wsController := websocketapi.GetWsHandler()
	go websocketapi.StartBrodcast() // loop que escuta as novas mensagens

	server.GET("/ws", gin.WrapH(wsController))

	go func() {
		err_ws := http.ListenAndServe(":8000", wsController)
		if err_ws != nil {
			panic("Error starting WebSocket server: " + err_ws.Error())
		}
	}()

	server.Run() // default 8080
}
