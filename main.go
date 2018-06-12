package main

import (
	"git.finogeeks.club/finochat/go-gin/api"
	"git.finogeeks.club/finochat/go-gin/config"
	"git.finogeeks.club/finochat/go-gin/logger"
	"git.finogeeks.club/finochat/go-gin/middleware"
	"github.com/gin-gonic/gin"
)

var version = "unknown"

func new() *gin.Engine {
	r := gin.New()
	r.Use(middleware.FinoMonitor())
	r.Use(gin.Logger(), gin.Recovery())

	baseGroup := r.Group("/api/v1/", func(c *gin.Context) {})
	baseGroup.GET("/item/:ID", api.GetHandler)
	baseGroup.POST("/item/", api.PostHandler)
	baseGroup.PUT("/item/:ID", api.PutHandler)
	baseGroup.DELETE("/item/:ID", api.DeleteHandler)

	return r

}

func main() {
	cfg := config.GetConfig()
	gin.SetMode(gin.ReleaseMode)
	r := new()
	var log = logger.GetLogger()
	log.Noticef("[gin-demo running... version %s]\n", version)
	r.Run(":" + cfg.HttpPort)
}
