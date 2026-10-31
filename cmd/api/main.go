package main

import (
	"log"
	"net/http"

	"github.com/codieSam/url-shortener-api/internal/handlers"
	"github.com/codieSam/url-shortener-api/internal/services"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	urlService := services.NewURLService()
	urlHandler := handlers.NewURLHandler(urlService)

	router.POST("/api/v1/urls", urlHandler.CreateURL)
	router.GET("./:code", urlHandler.RedirectURL)

	router.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"Service": "url-shortner-api",
		})
	})

	if err := router.Run(":8081"); err != nil {
		log.Fatal("Failed to start server", err)
	}

}
