package product

import (
	"context"
	"database/sql"
)

type RepositoryNative struct {
	db *sql.DB
}

func NewRepositoryNative(db *sql.DB) RepositoryNative {
	return RepositoryNative{
		db: db,
	}
}

func (r RepositoryNative) Create(ctx context.Context, req Product) (err error) {
	return
}
