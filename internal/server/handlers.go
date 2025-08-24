package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetQuotes(c *gin.Context) {
	symbol := c.DefaultQuery("symbols", "META")
	endpointUrl := fmt.Sprintf("%s/stocks/quotes/latest", ALPACA_MARKET_URL)

	resp, err := Client.R().
		SetQueryParam("symbols", symbol).
		Get(endpointUrl)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Request to /get-quotes failed"})
		return
	}

	c.Data(resp.StatusCode(), "application/json", resp.Body())
}

func PlaceOrder(c *gin.Context) {
	var orderReq struct {
		Symbol      string `json:"symbol"`
		Type        string `json:"type"`
		TimeInForce string `json:"time_in_force"`
		Qty         string `json:"qty"`
		Side        string `json:"side"`
	}

	if err := c.ShouldBindJSON(&orderReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	endpointUrl := fmt.Sprintf("%s/orders", ALPACA_TRADING_URL)
	resp, err := Client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(orderReq).
		Post(endpointUrl)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Request to /place-order failed"})
		return
	}

	c.Data(resp.StatusCode(), "application/json", resp.Body())
}

func GetAllOrders(c *gin.Context) {
	endpointUrl := fmt.Sprintf("%s/orders", ALPACA_TRADING_URL)

	resp, err := Client.R().Get(endpointUrl)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Request to /portfolio failed"})
		return
	}
	c.Data(resp.StatusCode(), "application/json", resp.Body())
}

func GetOpenPositions(c *gin.Context) {
	endpointUrl := fmt.Sprintf("%s/positions", ALPACA_TRADING_URL)

	resp, err := Client.R().
		Get(endpointUrl)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Request to /positions failed"})
		return
	}

	c.Data(resp.StatusCode(), "application/json", resp.Body())
}

func GetBars(c *gin.Context) {
	symbol := c.DefaultQuery("symbols", "META")
	timeframe := c.DefaultQuery("timeframe", "8Hour")

	endpointUrl := fmt.Sprintf("%s/stocks/bars", ALPACA_MARKET_URL)
	resp, err := Client.R().
		SetQueryParam("symbols", symbol).
		SetQueryParam("timeframe", timeframe).
		Get(endpointUrl)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Request to /bars failed"})
		return
	}

	c.Data(resp.StatusCode(), "application/json", resp.Body())
}
