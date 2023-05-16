package repository

import "github.com/jmoiron/sqlx"

type Repository struct {
	Store
}

func New(db *sqlx.DB) *Repository {
	return &Repository{
		Store:  NewPostgres(db),
	}
}
