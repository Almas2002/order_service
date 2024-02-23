package address_repository

import "github.com/jmoiron/sqlx"

type repository struct {
	table string
	db    *sqlx.DB
}

func New(db *sqlx.DB) *repository {
	return &repository{table: "address_repository", db: db}
}
