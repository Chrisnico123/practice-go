package service

import (
	"movies/models"
	"movies/repository"

	"github.com/labstack/echo/v4"
)

type MoviesSvc struct {
	repo repository.MoviesRepo
}

func NewMoviesSvc(repo repository.MoviesRepo) MoviesSvc {
	return MoviesSvc{
		repo: repo,
	}
}

func (svc MoviesSvc) CreateMovie(c echo.Context, movies models.Movies) (error) {
	err := svc.repo.CreateData(c.Request().Context() , movies)
	if err != nil {
		return err
	}	

	return nil
}

func (svc MoviesSvc) FindAllMovies(c echo.Context) ([]models.Movies , error) {
	var movies []models.Movies

	movies , err := svc.repo.GetAll(c.Request().Context())
	if err != nil {
		return nil , err
	}

	return movies , err
}

func (svc MoviesSvc) GetDataById(c echo.Context , id int) (models.Movies , error) {
	movies, err := svc.repo.GetDataById(c.Request().Context() , id)
	if err != nil {
		return models.Movies{} , err
	}

	return movies , nil
}
func (svc MoviesSvc) UpdateDataById(c echo.Context ,movie models.Movies, id int) (error) {

	err := svc.repo.UpdateDataById(c.Request().Context(), movie , id)

	if err != nil {
		return err
	}

	return nil
}
func (svc MoviesSvc) DeleteDataById(c echo.Context , id int) (error) {
	err := svc.repo.DeleteMovies(c.Request().Context() , id)

	if err != nil {
		return err
	}

	return nil
}
