package server

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/glebpepega/proj/internal/decoder"
	"github.com/glebpepega/proj/internal/person"
)

func (s *server) handlePerson(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		s.logger.Info("handle GET /people")
		ctx := context.Background()
		data, err := s.cache.Cache.Get(ctx, "people").Result()
		if err != nil {
			s.logger.Info("cache miss")
			if people, err := s.postgres.GetUsers(); err != nil {
				s.logger.Fatal(err)
			} else {
				pJson, err := json.Marshal(people)
				if err != nil {
					s.logger.Fatal(err)
				}
				data = `{"people":` + string(pJson) + `}`
				if err := s.cache.Cache.Set(ctx, "people", data, 0).Err(); err != nil {
					s.logger.Fatal(err)
				}
			}
		} else {
			s.logger.Info("cache hit")
		}
		if _, err := w.Write([]byte(data)); err != nil {
			s.logger.Fatal(err)
		}
		s.logger.Info("data retrieved")

	case "POST":
		s.logger.Info("handle POST /people")
		p := &person.Person{}
		if err := decoder.DecodeFromJSON(r.Body, &p); err != nil {
			s.logger.Info("request unsuccessful")
			http.Error(w, "could not unmarshal json",
				http.StatusBadRequest)
			return
		}
		if p.Name == "" || p.Surname == "" {
			s.logger.Info("request unsuccessful")
			http.Error(w, "name and surname required",
				http.StatusBadRequest)
			return
		}
		if err := s.postgres.StoreInDB(p); err != nil {
			s.logger.Fatal(err)
		} else {
			s.logger.Info("added id: ", p.ID)
		}

	case "PUT":
		s.logger.Info("handle PUT /people")
		p := &person.Person{}
		if err := decoder.DecodeFromJSON(r.Body, &p); err != nil {
			s.logger.Info("request unsuccessful")
			http.Error(w, "could not unmarshal json",
				http.StatusBadRequest)
			return
		}
		if p.Name == "" || p.Surname == "" {
			s.logger.Info("request unsuccessful")
			http.Error(w, "name and surname required",
				http.StatusBadRequest)
			return
		}
		if err := s.postgres.UpdateUser(p); err != nil {
			s.logger.Fatal(err)
		} else {
			s.logger.Info("updated id: ", p.ID)
		}

	case "DELETE":
		s.logger.Info("handle DELETE /people")
		p := person.Person{}
		if err := decoder.DecodeFromJSON(r.Body, &p); err != nil {
			s.logger.Info("request unsuccessful")
			http.Error(w, "could not unmarshal json",
				http.StatusBadRequest)
			return
		}
		if p.ID <= 0 {
			s.logger.Info("request unsuccessful")
			http.Error(w, "have to have an id that is greater than 0",
				http.StatusBadRequest)
			return
		}
		if err := s.postgres.DeleteUser(p.ID); err != nil {
			s.logger.Fatal(err)
		} else {
			s.logger.Info("deleted id: ", p.ID)
		}
	}
}
