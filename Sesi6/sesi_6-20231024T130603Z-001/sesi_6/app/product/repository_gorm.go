package product

import (
	"context"

	"gorm.io/gorm"
)

type PostgresGormRepository struct {
	db *gorm.DB
}

func NewPostgresGormRepository(db *gorm.DB) PostgresGormRepository {
	return PostgresGormRepository{
		db: db,
	}	
}

func (p PostgresGormRepository) Create(ctx context.Context, model Product) (err error) {
	return p.db.Create(&model).Error
}
