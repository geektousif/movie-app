package main

import (
	"github.com/geektousif/go-vue-movie-app/internal/config"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	cfg := config.Load()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"app":     cfg.AppName,
			"version": cfg.AppVersion,
		})
	})

	router.Run("0.0.0.0:" + cfg.Port)
}
