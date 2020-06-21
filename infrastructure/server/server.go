package server

import (
	"jiroshin/clean-architecture-golang/infrastructure/server/router"
	"log"
	"net/http"
)

func StartServe() {
	s := &http.Server{
		Addr:    ":8080",
		Handler: router.NewRouter(),
	}

	log.Println("Server started...")
	log.Printf("Listening on %s", s.Addr)
	if err := s.ListenAndServe(); err != nil {
		log.Printf("Failed starting server: %s", err)
		panic(err)
	}
}
