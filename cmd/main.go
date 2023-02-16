package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"main.go/cmd/pkg/db"
	"main.go/cmd/pkg/handlers"
)

func main() {
	newDB := db.InitializeDB()
	h := handlers.NewDB(newDB)
	router := mux.NewRouter()

	router.HandleFunc("/books", h.GetAllBooks).Methods("GET")
	router.HandleFunc("/books", h.AddBook).Methods("POST")
	router.HandleFunc("/books/{id}", h.GetBookByID).Methods("GET")
	router.HandleFunc("/books/{id}", h.UpdateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", h.DeleteBook).Methods("DELETE")

	fmt.Println("Server listening on port 9000")
	http.ListenAndServe(":9000", router)

}

// Tools like TablePlus can connect multiple databases