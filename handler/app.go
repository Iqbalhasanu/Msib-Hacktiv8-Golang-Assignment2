package handler

import (
	"github.com/Iqbalhasanu/Msib-Hacktiv8-Golang-Assignment2.git/database"
	_ "github.com/Iqbalhasanu/Msib-Hacktiv8-Golang-Assignment2.git/docs"
	"github.com/Iqbalhasanu/Msib-Hacktiv8-Golang-Assignment2.git/repository/item_repository/item_pg"
	"github.com/Iqbalhasanu/Msib-Hacktiv8-Golang-Assignment2.git/repository/order_repository/order_pg"
	"github.com/Iqbalhasanu/Msib-Hacktiv8-Golang-Assignment2.git/service"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	database.Initialize()
	db := database.GetDatabaseInstance()

	itemRepo := item_pg.NewItemPg(db)
	itemService := service.NewItemService(itemRepo)

	orderRepo := order_pg.NewOrderPg(db)
	orderService := service.NewOrderService(orderRepo, itemService)

	orderHandler := NewOrderHandler(orderService)

	route := gin.Default()

	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	orderRoute := route.Group("/orders")
	{
		orderRoute.POST("/", orderHandler.CreateOrder)
		orderRoute.GET("/", orderHandler.GetAllOrders)
		orderRoute.GET("/:orderId", orderHandler.GetOrder)
		orderRoute.PUT("/:orderId", orderHandler.UpdateOrder)
		orderRoute.DELETE("/:orderId", orderHandler.DeleteOrder)
	}

	route.Run(":8080")
}
