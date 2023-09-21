package server

import "github.com/glebpepega/proj/person"

func (s *server) StoreInDB(p *person.Person) error {
	if _, err := s.db.Exec("INSERT INTO posts (first_name,surname,patronymic,age,gender,country) VALUES ($1, $2, $3, $4, $5, $6)",
		p.Name, p.Surname, p.Patronymic, p.Age, p.Gender, p.Country); err != nil {
		return err
	}
	return nil
}
