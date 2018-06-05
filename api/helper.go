package api

import (
	"git.finogeeks.club/finochat/go-gin/logger"
	"github.com/gin-gonic/gin"
)

const SERVICE_NAME = "gin-demo"

var log = logger.GetLogger()

func newError(errcode string, err string) gin.H {
	return gin.H{"errcode": errcode, "error": err, "service": SERVICE_NAME}
}
