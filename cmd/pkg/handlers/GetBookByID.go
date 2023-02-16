package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"main.go/cmd/pkg/models"
)

func (h handler) GetBookByID(res http.ResponseWriter, req *http.Request) {
	var book models.Book
	params := mux.Vars(req)
	idx, _ := strconv.Atoi(params["id"])

	err := h.DB.Where("id : ?").First(&book, idx).Error
	HandleErr(err)

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(book)

}
