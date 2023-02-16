package handlers

import (
	"encoding/json"
	"main.go/cmd/pkg/models"
	"net/http"
)

func (h handler) GetAllBooks(res http.ResponseWriter, req *http.Request) {
	var book []models.Book
	err := h.DB.Find(&book).Error
	HandleErr(err)
	res.Header().Set("Content-Types", "application/json")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(&book)
}
