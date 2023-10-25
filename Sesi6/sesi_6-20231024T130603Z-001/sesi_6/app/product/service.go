package product

import (
	"context"
	"log"
)

type Repository interface {
	Create(ctx context.Context, req Product) (err error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return Service{
		repo: repo,
	}
}

func (s Service) CreateProduct(ctx context.Context, req Product) (err error) {
	if err = req.Validate(); err != nil {
		log.Println("error when try to validate request with error :", err.Error())
		return
	}

	if err = s.repo.Create(ctx, req); err != nil {
		log.Println("error when try to Create to database with error :", err.Error())
		return

	}

	return
}
