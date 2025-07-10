package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := gin.Default()
	router.Use(cors.Default())
	router.GET("/quote", getQuote)

	err = router.Run(":8080")
	if err != nil {
		panic(err)
	}
}

func getQuote(c *gin.Context) {
	symbol := c.DefaultQuery("symbol", "")
	c.JSON(200, gin.H {
		"symbol" : symbol,
		"price": 100.0,
	})
}
