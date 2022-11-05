package main

import (
	"net/http"
	"url-shortener/handler"
	"url-shortener/services"

	"github.com/gorilla/mux"
)

func main() {
	store := services.New()
	handler := handler.New(store)

	r := mux.NewRouter()

	r.HandleFunc("/shortlink", handler.GetShortURL).Methods(http.MethodPost)

	http.ListenAndServe(":8880", r)
}
