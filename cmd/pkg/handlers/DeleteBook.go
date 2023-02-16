package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"main.go/cmd/pkg/models"
)

func (h handler) DeleteBook(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	idx, _ := strconv.Atoi(params["id"])
	var book models.Book
	// for index, book := range mocks.BooksList {
	// 	if book.ID == idx {
	// 		mocks.BooksList = append(mocks.BooksList[:index], mocks.BooksList[index+1:]...)

	// 		break
	// 	}
	// }
	err := h.DB.Where("id : ?").First(&book, idx).Error
	HandleErr(err)
	h.DB.Delete(&book)
	json.NewEncoder(res).Encode(book)

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
}

// Some edge cases
// Deleting a resource that doesn't exist
// Updating a resource that doesn't exist
