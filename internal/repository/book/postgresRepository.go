package book

import "database/sql"

type postgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) postgresRepository {
	return postgresRepository{db: db}
}

func (p postgresRepository) Create() {
	//TODO implement me
	panic("implement me")
}

func (p postgresRepository) List() {
	//TODO implement me
	panic("implement me")
}

func (p postgresRepository) Get() {
	//TODO implement me
	panic("implement me")
}

func (p postgresRepository) Update() {
	//TODO implement me
	panic("implement me")
}

func (p postgresRepository) Delete() {
	//TODO implement me
	panic("implement me")
}
