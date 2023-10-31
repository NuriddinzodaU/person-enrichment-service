package repository

import (
	"person-service/internal/models"
)

func (r repository) GetPerson(id int64) (person *models.Person, err error) {
	db := r.Postgres.GetPostgresConnection()
	sqlQuery := `select p.name,p.surname, p.patronymic, p.age, p.gender, p.nationality from people p where id = ?;`
	if err := db.Raw(sqlQuery, id).Scan(&person).Error; err != nil {
		return
	}
	return
}
