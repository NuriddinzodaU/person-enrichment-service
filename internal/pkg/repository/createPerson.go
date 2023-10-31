package repository

import (
	"person-service/internal/models"
)

func (r repository) CreatePerson(person *models.Person) error {
	db := r.Postgres.GetPostgresConnection()
	sqlQuery := `insert into people (name, surname, patronymic, age, gender, nationality) values (?, ?, ?, ?, ?, ?);`
	if err := db.Exec(sqlQuery, person.Name, person.Surname, person.Patronymic, person.Age, person.Gender, person.Nationality).Error; err != nil {
		return err
	}
	return nil
}
