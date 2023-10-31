package repository

import (
	"person-service/internal/models"
)

func (r repository) GetAllPerson(count, page int) (person []*models.Person, err error) {
	db := r.Postgres.GetPostgresConnection()
	sqlQuery := `select p.name,p.surname, p.patronymic, p.age, p.gender, p.nationality from people p where limit ? offset ?;`
	if err := db.Raw(sqlQuery, count, page).Scan(&person).Error; err != nil {
		return nil, nil
	}
	return
}
