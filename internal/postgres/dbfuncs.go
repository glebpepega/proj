package postgres

import "github.com/glebpepega/proj/internal/person"

func (db *Postgres) StoreInDB(p *person.Person) error {
	if _, err := db.Postgres.Exec("INSERT INTO people (first_name, surname, patronymic, age, gender, country) VALUES ($1, $2, $3, $4, $5, $6)",
		p.Name, p.Surname, p.Patronymic, p.Age, p.Gender, p.OriginCountry); err != nil {
		return err
	}
	return nil
}

func (db *Postgres) GetUsers() (person.People, error) {
	people := person.People{}
	rows, err := db.Postgres.Query("SELECT * FROM people ORDER BY id;")
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

func (db *Postgres) UpdateUser(p *person.Person) error {
	if _, err := db.Postgres.Exec("UPDATE people SET first_name = $1, surname = $2, patronymic = $3, age = $4, gender = $5, country = $6 WHERE id = $7;",
		p.Name, p.Surname, p.Patronymic, p.Age, p.Gender, p.OriginCountry, p.ID); err != nil {
		return err
	}
	return nil
}

func (db *Postgres) DeleteUser(id int) error {
	if _, err := db.Postgres.Exec("DELETE FROM people WHERE id = $1;", id); err != nil {
		return err
	}
	return nil
}
