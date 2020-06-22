package api

import (
	"database/sql"
	"jiroshin/clean-architecture-golang/adapter/api/router"
	"jiroshin/clean-architecture-golang/infrastructure/database"
	"log"
	"net/http"
	"os"
)

func StartServe() {
	var db *sql.DB
	var err error

	if os.Getenv("DB_TYPE") == "mysql" {

		db, err = database.ConnectToMysql()
		defer db.Close()
		if err != nil {
			log.Printf("DB connection failure: %s", err)
			panic(err)
		}

	} else {

		user := os.Getenv("DB_USER")
		password := os.Getenv("DB_PASSWORD")
		dbName := os.Getenv("DB_NAME")
		host := os.Getenv("DB_HOST")
		port := os.Getenv("DB_PORT")

		db, err = database.ConnectToPostgres(host, port, user, password, dbName)
		defer db.Close()
		if err != nil {
			log.Printf("DB connection failure: %s", err)
			panic(err)
		}

	}

	s := &http.Server{
		Addr:    ":8080",
		Handler: router.NewRouter(db),
	}

	log.Println("Server started...")
	log.Printf("Listening on %s", s.Addr)
	if err := s.ListenAndServe(); err != nil {
		log.Printf("Failed starting server: %s", err)
		panic(err)
	}
}
