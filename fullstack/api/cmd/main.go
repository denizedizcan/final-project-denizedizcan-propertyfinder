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
	router.HandleFunc("/update-stocks", handlers.SetMiddlewareJSON(h.UpdateStock)).Methods("PUT")
	router.HandleFunc("/list-prices", handlers.SetMiddlewareJSON(h.GetPrice)).Methods("GET")
	router.HandleFunc("/add-prices", handlers.SetMiddlewareJSON(h.AddPrice)).Methods("PUT")
	router.HandleFunc("/get-basket", handlers.SetMiddlewareJSON(h.GetBasket)).Methods("GET")
	router.HandleFunc("/add-to-basket", handlers.SetMiddlewareJSON(h.AddBasketItem)).Methods("POST")
	router.HandleFunc("/update-basketitem", handlers.SetMiddlewareJSON(h.UpdateBasketItem)).Methods("POST")
	router.HandleFunc("/delete-basketitem", handlers.SetMiddlewareJSON(h.DeleteOneItem)).Methods("DELETE")
	router.HandleFunc("/add-order", handlers.SetMiddlewareJSON(h.AddOrder)).Methods("PUT")

	http.ListenAndServe(":12345", router)
}

// basket item add update test
