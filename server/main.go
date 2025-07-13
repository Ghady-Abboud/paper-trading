package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	//"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"github.com/joho/godotenv"
)

var client *resty.Client

func initResty() {
	client = resty.New().SetBaseURL(os.Getenv("ALPACA_MARKET_URL")).SetHeader("APCA-API-KEY-ID", os.Getenv("ALPACA_API_KEY")).SetHeader("APCA-API-SECRET-KEY",os.Getenv("ALPACA_SECRET_KEY"))
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	initResty()

	router := gin.Default()
	router.Use(cors.Default())
	router.GET("/get-quote", getQuote)

	err = router.Run(":8080")
	if err != nil {
		panic(err)
	}
}

func getQuote(c *gin.Context) {
	endpointUrl := fmt.Sprintf("%s/stocks/quotes/latest", os.Getenv("ALPACA_MARKET_URL"))

	resp, err := client.R().SetDebug(true).SetQueryParam("symbols", "AAPL").Get(endpointUrl)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch data"})
		return
	}

	c.Data(resp.StatusCode(), "application/json", resp.Body())
}
