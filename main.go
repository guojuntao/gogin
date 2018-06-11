package main

import (
	"net/http"
	_ "net/http/pprof"

	"git.finogeeks.club/finochat/go-gin/api"
	"git.finogeeks.club/finochat/go-gin/config"
	"git.finogeeks.club/finochat/go-gin/logger"
	"git.finogeeks.club/finochat/go-gin/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	// TODO: 条件编译，打 tag 的提交不编译 pprof
	go func() {
		// /debug/pprof
		http.ListenAndServe("0.0.0.0:8080", nil)
	}()
}

var VERSION string = "unknown"

func main() {
	cfg := config.GetConfig()
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(middleware.FinoMonitor())
	r.Use(gin.Logger(), gin.Recovery())

	baseGroup := r.Group("/api/v1/", func(c *gin.Context) {})
	baseGroup.GET("/item/:ID", api.GetHandler)
	baseGroup.POST("/item/", api.PostHandler)
	baseGroup.PUT("/item/:ID", api.PutHandler)
	baseGroup.DELETE("/item/:ID", api.DeleteHandler)

	var log = logger.GetLogger()
	log.Noticef("[gin-demo running... version %s]\n", VERSION)
	r.Run(":" + cfg.HttpPort)
}
