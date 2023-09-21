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
	db, err := sql.Open("postgres", "postgresql://test:1510@database:5432/test?sslmode=disable")
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
	s.db.Exec("CREATE TABLE wegreehtnfd (user_id serial PRIMARY KEY);")
	log.Fatal(http.ListenAndServe(":8180", s.router))
	s.Consume()
}
