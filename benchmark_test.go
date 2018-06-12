package main

import (
	"bytes"
	"encoding/json"
	"git.finogeeks.club/finochat/go-gin/api"
	"git.finogeeks.club/finochat/go-gin/db"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
)

// run ```go test -bench=.```
func BenchmarkGet(b *testing.B) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	baseGroup := r.Group("/api/v1/", func(c *gin.Context) {})
	baseGroup.GET("/item/:ID", api.GetHandler)
	baseGroup.PUT("/item/:ID", api.PutHandler)

	// PUT
	{
		reqItem := db.Item{
			ID:   "002",
			Name: "guojuntao",
			Age:  18,
		}
		reqBody, _ := json.Marshal(reqItem)

		req := httptest.NewRequest(http.MethodPut, "/api/v1/item/002", bytes.NewReader(reqBody))
		resp := httptest.NewRecorder()
		r.ServeHTTP(resp, req)

		if status := resp.Code; status != http.StatusOK {
			panic(status)
		}
	}

	// benchmark
	for i := 0; i < b.N; i++ {
		req := httptest.NewRequest(http.MethodGet, "/api/v1/item/002", nil)
		resp := httptest.NewRecorder()
		r.ServeHTTP(resp, req)
	}
}
