package server

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

type server struct {
	router *http.ServeMux
	db     *sql.DB
}

func initdb() *sql.DB {
	db, err := sql.Open("postgres", "postgresql://gleb:1510@database:5432/proj?sslmode=disable")
	if err != nil {
		fmt.Println(err)
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
	go s.Consume()
	s.router.HandleFunc("/person", s.handlePerson)
	log.Fatal(http.ListenAndServe(":8080", s.router))
}
