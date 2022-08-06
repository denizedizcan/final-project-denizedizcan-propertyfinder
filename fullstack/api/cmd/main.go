package main

import (
	"fmt"
	"net/http"

	"github.com/denizedizcan/final-project-denizedizcan-propertyfinder/api/db"
	"github.com/denizedizcan/final-project-denizedizcan-propertyfinder/api/handlers"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Starting App..")
	DB := db.Init()
	h := handlers.New(DB)
	router := mux.NewRouter()
	router.HandleFunc("/create-user", handlers.SetMiddlewareJSON(h.CreateUser)).Methods("POST")
	router.HandleFunc("/login", handlers.SetMiddlewareJSON(h.Login)).Methods("POST")

	http.ListenAndServe(":12345", router)
}
