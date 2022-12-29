package httpserver

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (d *Dependency) GetMovie(c echo.Context) error {
	movie, err := d.Movie.GetMovie(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Internal server error",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, movie)
}

func (d *Dependency) GetMovieByID(c echo.Context) error {
	id := c.Param("id")

	if id == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "movie id is required",
		})
	}

	movieId, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "movie id is invalid",
			"error":   err.Error(),
		})
	}

	movie, err := d.Movie.GetMovieById(c.Request().Context(), movieId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Internal server error",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, movie)
}
