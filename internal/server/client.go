package server

import (
	"log"
	"os"

	"github.com/go-resty/resty/v2"
	"github.com/joho/godotenv"
)

var Client *resty.Client
var ALPACA_API_KEY string
var ALPACA_SECRET_KEY string
var ALPACA_MARKET_URL string
var ALPACA_TRADING_URL string
var ALPACA_MARKET_WEBSOCKET_URL string
var ALPACA_TRADING_WEBSOCKET_URL string

func RestyClientInit() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	ALPACA_API_KEY = os.Getenv("ALPACA_API_KEY")
	ALPACA_SECRET_KEY = os.Getenv("ALPACA_SECRET_KEY")
	ALPACA_MARKET_URL = os.Getenv("ALPACA_MARKET_URL")
	ALPACA_TRADING_URL = os.Getenv("ALPACA_TRADING_URL")
	ALPACA_MARKET_WEBSOCKET_URL = os.Getenv("ALPACA_MARKET_WEBSOCKET_URL")
	ALPACA_TRADING_WEBSOCKET_URL = os.Getenv("ALPACA_TRADING_WEBSOCKET_URL")

	Client = resty.New().
		SetHeader("APCA-API-KEY-ID", ALPACA_API_KEY).
		SetHeader("APCA-API-SECRET-KEY", ALPACA_SECRET_KEY)
}
