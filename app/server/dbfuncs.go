package server

import (
	"github.com/glebpepega/proj/person"
)

func (s *server) storeInDB(p *person.Person) error {
	if _, err := s.db.Exec("INSERT INTO people (first_name, surname, patronymic, age, gender, country) VALUES ($1, $2, $3, $4, $5, $6)",
		p.Name, p.Surname, p.Patronymic, p.Age, p.Gender, p.OriginCountry); err != nil {
		return err
	}
	return nil
}

func (s *server) getUsers(pag Pagination) (person.People, error) {
	people := person.People{}
	offset := pag.PageNum*pag.PageSize - 1
	rows, err := s.db.Query("SELECT * FROM people ORDER BY id LIMIT $1 OFFSET $2;", pag.PageSize, offset)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		person := person.Person{}
		if err := rows.Scan(&person.ID, &person.Name, &person.Surname, &person.Patronymic, &person.Age, &person.Gender, &person.OriginCountry); err != nil {
			return nil, err
		}
		people = append(people, person)
	}
	return people, nil
}

func (s *server) updateUser(p *person.Person) error {
	if _, err := s.db.Exec("UPDATE people SET first_name = $1, surname = $2, patronymic = $3, age = $4, gender = $5, country = $6 WHERE id = $7;",
		p.Name, p.Surname, p.Patronymic, p.Age, p.Gender, p.OriginCountry, p.ID); err != nil {
		return err
	}
	return nil
}

func (s *server) deleteUser(id int) error {
	if _, err := s.db.Exec("DELETE FROM people WHERE id = $1;", id); err != nil {
		return err
	}
	return nil
}
