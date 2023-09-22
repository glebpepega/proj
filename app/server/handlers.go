package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/glebpepega/proj/decoder"
	"github.com/glebpepega/proj/person"
	"github.com/go-playground/validator/v10"
)

type Pagination struct {
	PageSize int `validate:"required,gt=0"`
	PageNum  int `validate:"required,gt=0"`
}

type id struct {
	ID int `validate:"required,gt=0"`
}

func (s *server) handlePerson(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		pag := &Pagination{}
		if err := decoder.DecodeFromJSON(r.Body, &pag); err != nil {
			http.Error(w, "could not unmarshal json",
				http.StatusBadRequest)
			return
		}
		validate := validator.New()
		if err := validate.Struct(pag); err != nil {
			http.Error(w, "both pagenum and pagesize have to be present and greater than 0",
				http.StatusBadRequest)
			return
		}
		if people, err := s.getUsers(*pag); err != nil {
			log.Fatal(err)
		} else {
			pJson, err := json.Marshal(people)
			if err != nil {
				log.Fatal(err)
			}
			w.Write([]byte(`{"people":` + string(pJson) + `}`))
		}

	case "POST":
		p := &person.Person{}
		if err := decoder.DecodeFromJSON(r.Body, &p); err != nil {
			http.Error(w, "could not unmarshal json",
				http.StatusBadRequest)
			return
		}
		validate := validator.New()
		if err := validate.Struct(p); err != nil {
			http.Error(w, "name and surname required",
				http.StatusBadRequest)
			return
		}
		if err := s.storeInDB(p); err != nil {
			log.Fatal(err)
		}

	case "PUT":
		p := &person.Person{}
		if err := decoder.DecodeFromJSON(r.Body, &p); err != nil {
			http.Error(w, "could not unmarshal json",
				http.StatusBadRequest)
			return
		}
		validate := validator.New()
		if err := validate.Struct(p); err != nil {
			http.Error(w, "name and surname required",
				http.StatusBadRequest)
			return
		}
		if err := s.updateUser(p); err != nil {
			log.Fatal(err)
		}

	case "DELETE":
		i := &id{}
		if err := decoder.DecodeFromJSON(r.Body, &i); err != nil {
			http.Error(w, "could not unmarshal json",
				http.StatusBadRequest)
			return
		}
		validate := validator.New()
		if err := validate.Struct(i); err != nil {
			http.Error(w, "have to have an id that is greater than 0",
				http.StatusBadRequest)
			return
		}
		if err := s.deleteUser(i.ID); err != nil {
			log.Fatal(err)
		}
	}
}
