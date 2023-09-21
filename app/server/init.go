package server

import (
	"database/sql"
	"log"
	"net/http"
)

type server struct {
	router *http.ServeMux
	db     *sql.DB
}

func initdb() *sql.DB {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=gleb password=1510 dbname=proj sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	return db
}

func New() *server {
	return &server{
		router: http.NewServeMux(),
		db:     initdb(),
	}
}

func (s *server) Start() {
	go log.Fatal(http.ListenAndServe(":8080", s.router))
	s.Consume()
}
