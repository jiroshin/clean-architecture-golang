package handler

import (
	"net/http"
)

func CreateBook() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("200 OK"))
	}
}
