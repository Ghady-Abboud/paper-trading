package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"github.com/joho/godotenv"
)

var client *resty.Client
var ALPACA_MARKET_URL string
var ALPACA_API_KEY string
var ALPACA_SECRET_KEY string

func initResty() {
	client = resty.New().
		SetBaseURL(ALPACA_MARKET_URL).
		SetHeader("APCA-API-KEY-ID", ALPACA_API_KEY).
		SetHeader("APCA-API-SECRET-KEY",ALPACA_SECRET_KEY)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ALPACA_MARKET_URL = os.Getenv("ALPACA_MARKET_URL")
	ALPACA_API_KEY = os.Getenv("ALPACA_API_KEY")
	ALPACA_SECRET_KEY = os.Getenv("ALPACA_SECRET_KEY")

	initResty()

	router := gin.Default()
	router.Use(cors.Default())
	router.GET("/default-quotes", defaultQuotes)

	err = router.Run(":8080")
	if err != nil {
		panic(err)
	}
}

func defaultQuotes(c *gin.Context) {
	symbols := [5]string {"AAPL", "AMZN", "TSLA", "GOOG", "META"}
	joinedSymbols := strings.Join(symbols[:], ",")
	endpointUrl := fmt.Sprintf("%s/stocks/quotes/latest", ALPACA_MARKET_URL)

	resp, err := client.R().SetDebug(true).SetQueryParam("symbols", joinedSymbols).Get(endpointUrl)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch data"})
		return
	}

	c.Data(resp.StatusCode(), "application/json", resp.Body())
}
