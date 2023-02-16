package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"main.go/cmd/pkg/models"
	"net/http"
	"strconv"
)

func (h handler) UpdateBook(res http.ResponseWriter, req *http.Request) {
	var updatedBooks models.Book
	defer req.Body.Close()
	params := mux.Vars(req)
	idx, _ := strconv.Atoi(params["id"])
	body, err := ioutil.ReadAll(req.Body)
	json.Unmarshal(body, &updatedBooks)
	HandleErr(err)
	var book models.Book

	err = h.DB.Where("id : ?").First(&book, idx).Error
	HandleErr(err)
	book.Title = updatedBooks.Title
	book.Desc = updatedBooks.Desc
	book.Author = updatedBooks.Author

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(book)

}
