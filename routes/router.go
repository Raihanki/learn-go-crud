package routes

import (
	"crud_app/controllers/author"
	"crud_app/controllers/book"

	"github.com/gorilla/mux"
)

func Router(route *mux.Router) {
	api := route.PathPrefix("/api").Subrouter()

	//author
	authorRoute := api.PathPrefix("/authors").Subrouter()
	authorRoute.HandleFunc("", author.Index).Methods("GET")
	authorRoute.HandleFunc("", author.Store).Methods("POST")
	authorRoute.HandleFunc("/{id}", author.Show).Methods("GET")
	authorRoute.HandleFunc("/{id}", author.Update).Methods("PUT")
	authorRoute.HandleFunc("/{id}", author.Destroy).Methods("DELETE")

	//book
	bookRoute := api.PathPrefix("/books").Subrouter()
	bookRoute.HandleFunc("", book.Index).Methods("GET")
	bookRoute.HandleFunc("", book.Store).Methods("POST")
	bookRoute.HandleFunc("/{id}", book.Update).Methods("PUT")
	bookRoute.HandleFunc("/{id}", book.Show).Methods("GET")
	bookRoute.HandleFunc("/{id}", book.Destroy).Methods("DELETE")
}
