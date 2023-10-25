package handler

import (
	"movies/models"
	"movies/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type MovieHandler struct {
	svc service.MoviesSvc
}

func NewMovieHandler(svc service.MoviesSvc) MovieHandler {
	return MovieHandler{
		svc: svc,
	}
}

func (h MovieHandler) CreateDataMovie(c echo.Context) error {
	movie := new(models.Movies)
	c.Bind(movie)

	err := h.svc.CreateMovie(c,*movie)

	if err != nil {
		return c.JSON(http.StatusBadRequest , echo.Map{
			"success": false,
			"message": "ERR BAD REQUEST",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusAccepted , echo.Map{
		"success": true,
		"message": "CREATE SUCCESS",
	})
}

func (h MovieHandler) FindDataMovie(c echo.Context) error {
	var movies []models.Movies

	movies , err := h.svc.FindAllMovies(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest , echo.Map{
			"success": false,
			"message": "ERR BAD REQUEST",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusAccepted , echo.Map{
		"success": true,
		"message": "GET ALL SUCCESS",
		"Data" : movies,
	})
}


func (h MovieHandler) GetMoviesById(c echo.Context) error {
	var data models.Movies
	idParams := c.Param("id")
	id , err := strconv.Atoi(idParams)
	if err != nil {
		return c.JSON(http.StatusBadGateway , echo.Map{
			"success": false,
			"message": "ERR BAD REQUEST",
			"error":   err.Error(),
		})
	}

	data , err = h.svc.GetDataById(c , id)
	if err != nil {
		return c.JSON(http.StatusBadGateway , echo.Map{
			"success": false,
			"message": "ERR BAD REQUEST",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusAccepted , echo.Map{
		"success": true,
		"message" : "GET DATA SUCCESS",
		"payload" : data,
	})
}
func (h MovieHandler) UpdateMoviesById(c echo.Context) error {
	req := new(models.Movies)
	err := c.Bind(req)
	if err != nil {
		return c.JSON(http.StatusBadGateway , echo.Map{
			"success": false,
			"message": "ERR BAD REQUEST",
			"error":   err.Error(),
		})
	}


	idParams := c.Param("id")
	id , err := strconv.Atoi(idParams)
	if err != nil {
		return c.JSON(http.StatusBadGateway , echo.Map{
			"success": false,
			"message": "ERR BAD REQUEST",
			"error":   err.Error(),
		})
	}

	err = h.svc.UpdateDataById(c , *req , id)
	if err != nil {
		return c.JSON(http.StatusBadGateway , echo.Map{
			"success": false,
			"message": "ERR BAD REQUEST",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusAccepted , echo.Map{
		"success": true,
		"message" : "UPDATE DATA SUCCESS",
	})
}
func (h MovieHandler) DeleteMovies(c echo.Context) error {
	idParams := c.Param("id")
	id , err := strconv.Atoi(idParams)
	if err != nil {
		return c.JSON(http.StatusBadGateway , echo.Map{
			"success": false,
			"message": "ERR BAD REQUEST",
			"error":   err.Error(),
		})
	}
	err = h.svc.DeleteDataById(c, id)
	if err != nil {
		return c.JSON(http.StatusBadGateway , echo.Map{
			"success": false,
			"message": "ERR BAD REQUEST",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusAccepted , echo.Map{
		"success": true,
		"message" : "DELETE DATA SUCCESS",
	})	
}