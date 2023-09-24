package postgres

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

type Postgres struct {
	Postgres *sql.DB
}

func New() *Postgres {
	return &Postgres{}
}

func (p *Postgres) Configure() error {
	POSTGRES_USER := os.Getenv("POSTGRES_USER")
	POSTGRES_PASSWORD := os.Getenv("POSTGRES_PASSWORD")
	POSTGRES_DB := os.Getenv("POSTGRES_DB")
	dataSourceName := fmt.Sprintf("postgresql://%s:%s@database:5432/%s?sslmode=disable", POSTGRES_USER, POSTGRES_PASSWORD, POSTGRES_DB)
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}
	p.Postgres = db
	return nil
}
