package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"Service": "url-shortner-api",
		})
	})

	if err := router.Run(":8181"); err != nil {
		log.Fatal("Failed to start server", err)
	}

}
