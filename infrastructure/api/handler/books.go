package handler

import (
	"encoding/json"
	"jiroshin/clean-architecture-golang/entity"
	"jiroshin/clean-architecture-golang/usecase"
	"log"
	"net/http"
)

func CreateBook(uc *usecase.Book) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		requestBook := &entity.Book{}

		if err := json.NewDecoder(r.Body).Decode(requestBook); err != nil {
			log.Printf("Request: %s %s, unable to parse content: %v", r.Method, r.URL, err)
			http.Error(w, "500 internal server error", http.StatusInternalServerError)
			return
		}

		if err := uc.CreateBook(requestBook); err != nil {
			log.Printf("Request: %s %s, unable to parse content: %v", r.Method, r.URL, err)
			http.Error(w, "500 internal server error", http.StatusInternalServerError)
			return
		}

		bytes, err := json.Marshal(requestBook)
		if err != nil {
			log.Printf("Request: %s %s, unable to parse content: %v", r.Method, r.URL, err)
			http.Error(w, "500 internal server error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		if _, err = w.Write(bytes); err != nil {
			log.Printf("Request: %s %s, unable to parse content: %v", r.Method, r.URL, err)
		}
	}
}
