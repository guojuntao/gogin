package main

import (
	"git.finogeeks.club/finochat/go-gin/api"
	"git.finogeeks.club/finochat/go-gin/config"
	"git.finogeeks.club/finochat/go-gin/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.GetConfig()
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(middleware.FinoMonitor())
	router.Use(gin.Recovery())

	baseGroup := router.Group("/api/v1/", func(c *gin.Context) {})
	baseGroup.GET("/item/:ID", api.GetHandler)
	baseGroup.POST("/item/", api.PostHandler)
	baseGroup.PUT("/item/:ID", api.PutHandler)
	baseGroup.DELETE("/item/:ID", api.DeleteHandler)

	router.Run(":" + cfg.HttpPort)
}
