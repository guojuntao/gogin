package main

import (
	"bytes"
	"encoding/json"
	"git.finogeeks.club/finochat/go-gin/db"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestItemCRUD(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := new()

	reqItem := db.Item{
		ID:   "001",
		Name: "guojuntao",
		Age:  18,
	}
	reqBody, _ := json.Marshal(reqItem)

	newItem := db.Item{
		ID:   "001",
		Name: "guojuntao",
		Age:  24,
	}
	newBody, _ := json.Marshal(newItem)

	// 1. Post Item
	{
		req := httptest.NewRequest(http.MethodPost, "/api/v1/item/", bytes.NewReader(reqBody))
		resp := httptest.NewRecorder()
		r.ServeHTTP(resp, req)

		if status := resp.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}
	}

	// 2. Get Item
	{
		req := httptest.NewRequest(http.MethodGet, "/api/v1/item/001", nil)
		resp := httptest.NewRecorder()
		r.ServeHTTP(resp, req)

		if resp.Body.String() != string(reqBody) {
			t.Errorf("handler returned unexpected body: got %v want %v",
				resp.Body.String(), string(reqBody))
		}
	}

	// 3. Modify Item (PUT)
	{
		req := httptest.NewRequest(http.MethodPut, "/api/v1/item/001", bytes.NewReader(newBody))
		resp := httptest.NewRecorder()
		r.ServeHTTP(resp, req)

		if status := resp.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}
	}

	// 4. Get New Item
	{
		req := httptest.NewRequest(http.MethodGet, "/api/v1/item/001", nil)
		resp := httptest.NewRecorder()
		r.ServeHTTP(resp, req)

		if resp.Body.String() != string(newBody) {
			t.Errorf("handler returned unexpected body: got %v want %v",
				resp.Body.String(), string(newBody))
		}
	}

	// 5. Delete Item
	{
		req := httptest.NewRequest(http.MethodDelete, "/api/v1/item/001", nil)
		resp := httptest.NewRecorder()
		r.ServeHTTP(resp, req)

		if status := resp.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}
	}
}
