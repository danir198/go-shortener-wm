package main

import (
	"fmt"

	"github.com/danir198/go-shortener-wm/handler"
	"github.com/danir198/go-shortener-wm/store"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Printf("Hello Go URL Shortener !🚀")
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hey Go URL Shortener !",
		})
	})
	r.POST("/create-short-url", func(c *gin.Context) {
		handler.CreateShortUrl(c)
	})
	r.GET("/:shortUrl", func(c *gin.Context) {
		handler.HandleShortUrlRedirect(c)
	})

	store.InitializeStore()

	err := r.Run(":9808")
	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server - Error: %v", err))
	}
}
