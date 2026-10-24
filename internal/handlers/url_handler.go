package handlers

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/codieSam/url-shortener-api/internal/services"
	"github.com/gin-gonic/gin"
)

type URLHandler struct {
	service *services.URLService
}

func NewURLHandler(service *services.URLService) *URLHandler {
	return &URLHandler{
		service: service,
	}
}

type createURLRequest struct {
	URL string `json:"url" binding:"required"`
}

func (h *URLHandler) CreateURL(c *gin.Context) {
	var request createURLRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Request must contain a valid JSON body with url field.",
		})
		return
	}

	originalURL := strings.TrimSpace(request.URL)

	parsedURL, err := url.Parse(originalURL)

	if err != nil ||
		parsedURL.Hostname() == "" ||
		(parsedURL.Scheme != "http" && parsedURL.Scheme != "https") {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "URL must be an absolute HTTP or HTTPS URL.",
		})
		return

	}

	code, err := h.service.Create(originalURL)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create short URL.",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"short_code":   code,
		"original_url": originalURL,
	})
}
