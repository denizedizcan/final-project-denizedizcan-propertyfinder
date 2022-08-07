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
	router.HandleFunc("/list-products", handlers.SetMiddlewareJSON(h.GetProducts)).Methods("GET")
	router.HandleFunc("/add-products", handlers.SetMiddlewareJSON(h.AddProducts)).Methods("PUT")
	router.HandleFunc("/list-stocks", handlers.SetMiddlewareJSON(h.GetStocks)).Methods("GET")
	router.HandleFunc("/add-stocks", handlers.SetMiddlewareJSON(h.AddStocks)).Methods("PUT")
	router.HandleFunc("/list-prices", handlers.SetMiddlewareJSON(h.GetPrice)).Methods("GET")
	router.HandleFunc("/add-prices", handlers.SetMiddlewareJSON(h.AddPrice)).Methods("PUT")
	router.HandleFunc("/get-basket", handlers.SetMiddlewareJSON(h.GetBasket)).Methods("GET")
	router.HandleFunc("/add-to-basket", handlers.SetMiddlewareJSON(h.GetBasket)).Methods("POST")
	router.HandleFunc("/update-basket", handlers.SetMiddlewareJSON(h.GetBasket)).Methods("POST")

	http.ListenAndServe(":12345", router)
}
