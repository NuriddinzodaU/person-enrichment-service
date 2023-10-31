package repository

import (
	"person-service/internal/models"
)

func (r repository) UpdatePerson(person *models.Person) error {
	db := r.Postgres.GetPostgresConnection()
	sqlQuery := `update people set (name = ?,
    								surname = ?,
    								patronymic = ?,
    								age = ?,
    								gender = ?,
    								nationality = ?) 
              	  where id = ?;`
	if err := db.Exec(sqlQuery, person.Name, person.Surname, person.Patronymic, person.Age, person.Gender, person.Nationality, person.ID).Error; err != nil {
		return err
	}
	return nil
}
