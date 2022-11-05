package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"url-shortener/models"
	"url-shortener/services"

	"github.com/asaskevich/govalidator"
)

type urlHandler struct {
	dataStore services.URL
}

func New(datastore services.URL) *urlHandler {
	return &urlHandler{dataStore: datastore}
}

func (u *urlHandler) GetShortURL(w http.ResponseWriter, r *http.Request) {
	var url models.URL

	if err := json.NewDecoder(r.Body).Decode(&url); err != nil {
		writeAPIResponse(w, "bad request", http.StatusBadRequest)
		return
	}

	if !govalidator.IsURL(url.LongURL) {
		writeAPIResponse(w, "invalid url", http.StatusBadRequest)
		return
	}

	shortUrl, err := u.dataStore.GetShortURL(url)
	if err != nil {
		log.Println(err.Error())
		writeAPIResponse(w, "unexpected error occured", http.StatusInternalServerError)
		return
	}

	writeAPIResponse(w, shortUrl, http.StatusOK)
}

func writeAPIResponse(w http.ResponseWriter, msg string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write([]byte(msg))
}
