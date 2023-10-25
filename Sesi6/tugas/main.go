package main

import (
	"movies/config"
	"movies/handler"
	"movies/repository"
	"movies/service"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	router := echo.New()
	db, err := config.ConnectDb()
	if err != nil {
		panic(err)
	}

	repo := repository.NewMoviesRepository(db)
	svc := service.NewMoviesSvc(repo)
	handler := handler.NewMovieHandler(svc)
	

	router.POST("/movie", handler.CreateDataMovie)
	router.GET("/movie", handler.FindDataMovie)
	router.GET("/movie/:id", handler.GetMoviesById)
	router.PUT("/movie/:id", handler.UpdateMoviesById)
	router.DELETE("/movie/:id", handler.DeleteMovies)

	router.Start(":4000")
}