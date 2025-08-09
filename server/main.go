package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"github.com/joho/godotenv"
)

var client *resty.Client
var ALPACA_MARKET_URL string
var ALPACA_API_KEY string
var ALPACA_SECRET_KEY string
var ALPACA_TRADING_URL string

func initResty() {
	client = resty.New().
		SetHeader("APCA-API-KEY-ID", ALPACA_API_KEY).
		SetHeader("APCA-API-SECRET-KEY",ALPACA_SECRET_KEY)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ALPACA_MARKET_URL = os.Getenv("ALPACA_MARKET_URL")
	ALPACA_TRADING_URL = os.Getenv("ALPACA_TRADING_URL")
	ALPACA_API_KEY = os.Getenv("ALPACA_API_KEY")
	ALPACA_SECRET_KEY = os.Getenv("ALPACA_SECRET_KEY")

	initResty()

	router := gin.Default()
	router.Use(cors.Default())
	router.GET("/api/get-quotes", GetQuotes)
	router.POST("/api/place-order", PlaceOrder)
	router.GET("/api/get-orders", GetAllOrders)
	router.GET("/api/get-positions", GetOpenPositions)

	err = router.Run(":8080")
	if err != nil {
		panic(err)
	}
}

func GetQuotes(c *gin.Context) {
	symbol := c.DefaultQuery("symbols", "META")
	endpointUrl := fmt.Sprintf("%s/stocks/quotes/latest", ALPACA_MARKET_URL)

	resp, err := client.R().
							SetDebug(true).
							SetQueryParam("symbols", symbol).
							Get(endpointUrl)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Request to /default-quotes failed"})
		return
	}

	c.Data(resp.StatusCode(), "application/json", resp.Body())
}

func PlaceOrder(c *gin.Context) {
	var orderReq struct {
		Symbol string `json:"symbol"`
		Type string `json:"type"`
		TimeInForce string `json:"time_in_force"`
		Qty string `json:"qty"`
		Side string `json:"side"`
	}
	
	if err := c.ShouldBindJSON(&orderReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	endpointUrl := fmt.Sprintf("%s/orders", ALPACA_TRADING_URL)
	resp, err := client.R().
							SetHeader("Content-Type", "application/json").
							SetBody(orderReq).
							Post(endpointUrl)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error" : "Request to /place-order failed"})
		return
	}

	c.Data(resp.StatusCode(), "application/json", resp.Body())
}

func GetAllOrders(c *gin.Context) {
	endpointUrl := fmt.Sprintf("%s/orders", ALPACA_TRADING_URL)

	resp, err := client.R().Get(endpointUrl)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Request to /portfolio failed"})
		return
	}
	c.Data(resp.StatusCode(), "application/json", resp.Body())
}

func GetOpenPositions(c *gin.Context) {
	endpointUrl := fmt.Sprintf("%s/positions", ALPACA_TRADING_URL)

	resp, err := client.R().Get(endpointUrl)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error" : "Request to /positions failed"})
		return
	}

	c.Data(resp.StatusCode(), "application/json", resp.Body())
}