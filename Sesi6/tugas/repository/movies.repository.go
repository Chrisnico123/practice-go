package repository

import (
	"context"
	"movies/models"

	"gorm.io/gorm"
)

type MoviesRepo struct {
	db *gorm.DB
}

func NewMoviesRepository(db *gorm.DB) (MoviesRepo) {
	return MoviesRepo{
		db: db,
	}
}

func (m MoviesRepo) CreateData(c context.Context, movies models.Movies) (err error) {
	return m.db.Create(&movies).Error
}

func (m MoviesRepo) GetAll(c context.Context) ([]models.Movies , error) {
	var movies []models.Movies

	result := m.db.Find(&movies)

	if result.Error != nil {
		return nil , result.Error
	}

	return movies , nil
}

func (m MoviesRepo) GetDataById(c context.Context, id int) (models.Movies , error) {
	var movies models.Movies

	result := m.db.First(&movies , id)

	if result.Error != nil {
		return models.Movies{}, result.Error
	}

	return movies , nil
}

func (m MoviesRepo) UpdateDataById(c context.Context , movies models.Movies , id int) (error) {
	result := m.db.Model(&models.Movies{}).Where("id = ?", id).Updates(models.Movies{
		Title:       movies.Title,
		Description: movies.Description,
		Genre:      movies.Genre,
	})
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (m MoviesRepo) DeleteMovies(c context.Context , id int) (error) {
	result := m.db.Delete(models.Movies{} , id)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
