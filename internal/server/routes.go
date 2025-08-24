package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes() *gin.Engine{
	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/api/get-quotes", GetQuotes)
	router.POST("/api/place-order", PlaceOrder)
	router.GET("/api/get-orders", GetAllOrders)
	router.GET("/api/get-positions", GetOpenPositions)
	router.GET("/api/get-bars", GetBars)

	return router
}
