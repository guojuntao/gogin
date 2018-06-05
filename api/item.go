package api

import (
	"git.finogeeks.club/finochat/go-gin/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetHandler(c *gin.Context) {
	ID := c.Param("ID")

	item, err := db.FindItem(ID)
	if err != nil {
		log.Errorf("FindItem [%s] error [%v]", ID, err)
		c.JSON(http.StatusNotFound, newError("FC_INVALID_PARAMS", "查找失败"))
		return
	}

	c.JSON(http.StatusOK, item)
}

func PostHandler(c *gin.Context) {
	var reqItem db.Item
	if err := c.BindJSON(&reqItem); err != nil {
		log.Errorf("BindJSON error [%v]", err)
		c.JSON(http.StatusBadRequest, newError("FC_INVALID_PARAMS", "请求格式不合法"))
		return
	}

	if err := db.InsertItem(reqItem); err != nil {
		log.Errorf("InsertItem [%+v] error [%v]", reqItem, err)
		c.JSON(http.StatusNotFound, newError("FC_INVALID_PARAMS", "插入失败"))
		return
	}

	c.Status(http.StatusOK)
}

func PutHandler(c *gin.Context) {
	ID := c.Param("ID")

	var reqItem db.Item
	if err := c.BindJSON(&reqItem); err != nil {
		log.Errorf("BindJSON error [%v]", err)
		c.JSON(http.StatusBadRequest, newError("FC_INVALID_PARAMS", "请求格式不合法"))
		return
	}

	if err := db.UpdateItem(ID, reqItem); err != nil {
		log.Errorf("UpdateItem [%s][%+v] error [%v]", ID, reqItem, err)
		c.JSON(http.StatusBadRequest, newError("FC_INVALID_PARAMS", "更新失败"))
		return
	}

	c.Status(http.StatusOK)
}

func DeleteHandler(c *gin.Context) {
	ID := c.Param("ID")

	if err := db.DeleteItem(ID); err != nil {
		log.Errorf("DeleteItem [%s] error [%v]", ID, err)
		c.JSON(http.StatusBadRequest, newError("FC_INVALID_PARAMS", "删除失败"))
		return
	}

	c.Status(http.StatusOK)
}
