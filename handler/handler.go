package handler

import (
	"net/http"

	"github.com/danir198/go-shortener-wm/shortener"
	"github.com/danir198/go-shortener-wm/store"
	"github.com/gin-gonic/gin"
)

type UrlCreationRequest struct {
	LongUrl string `json:"long_url" binding:"required"`
	UserId  string `json:"user_id" binding:"required"`
}

func CreateShortUrl(c *gin.Context) {
	var creationRequest UrlCreationRequest
	if err := c.ShouldBindJSON(&creationRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	shortUrl := shortener.GeneratorShortLink(creationRequest.LongUrl, creationRequest.UserId)
	store.SaveUrlMapping(shortUrl, creationRequest.LongUrl, creationRequest.UserId)

	host := "http:/localhost:9808"
	c.JSON(200, gin.H{
		"message":   "short url succesfully created",
		"short_url": host + shortUrl,
	})

}

func HandleShortUrlRedirect(c *gin.Context) {

	shortUrl := c.Param("shortUrl")
	initialUrl := store.RetrieveInitialUrl(shortUrl)
	c.Redirect(302, initialUrl)
}
