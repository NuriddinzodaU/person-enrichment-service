package repository

func (r repository) DeletePerson(id int) error {
	db := r.Postgres.GetPostgresConnection()
	sqlQuery := `delete from people where id = ?;`
	if err := db.Exec(sqlQuery, id).Error; err != nil {
		return err
	}
	return nil
}
