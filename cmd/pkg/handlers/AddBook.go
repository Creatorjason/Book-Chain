package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	// "main.go/cmd/pkg/mocks"
	"main.go/cmd/pkg/models"
	// "math/rand"
	"net/http"
)

func HandleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func (h handler) AddBook(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
	defer req.Body.Close()
	body, err := ioutil.ReadAll(req.Body)
	HandleErr(err)

	var book models.Book

	// Decode json data to struct
	err = json.Unmarshal(body, &book)
	HandleErr(err)
	err = h.DB.Create(&book).Error
	if err != nil {
		// res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode("Error adding new block to DB")
	}
	//Manually assign id before appending

	// Send newly added book struct back to client
	json.NewEncoder(res).Encode(book)

}
