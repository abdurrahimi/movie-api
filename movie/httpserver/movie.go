package httpserver

import (
	"fmt"
	"movie-api/movie"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
)

type request struct {
	ID          int64      `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Rating      float64    `json:"rating"`
	Image       string     `json:"image"`
	CreatedAt   CustomTime `json:"created_at,omitempty"`
	UpdatedAt   CustomTime `json:"updated_at,omitempty"`
}

type CustomTime struct {
	time.Time
}

func (t CustomTime) MarshalJSON() ([]byte, error) {
	date := t.Time.Format("2006-01-02 15:04:05")
	date = fmt.Sprintf(`"%s"`, date)
	return []byte(date), nil
}

func (t *CustomTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")

	date, err := time.Parse("2006-01-02 15:04:05", s)
	if err != nil {
		return err
	}
	t.Time = date
	return
}

func (d *Dependency) GetMovie(c echo.Context) error {
	movie, err := d.Movie.GetMovie(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Internal server error",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Success",
		"movies":  movie,
	})
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

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Success",
		"movie":   movie,
	})
}

func (d *Dependency) CreateMovie(c echo.Context) error {
	var req request
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "invalid movie data",
			"error":   err.Error(),
		})
	}

	if req.Title == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "movie title is required",
		})
	}

	if req.Description == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "movie description is required",
		})
	}

	movie := &movie.Movies{
		Title:       req.Title,
		Description: req.Description,
		Rating:      req.Rating,
		Image:       req.Image,
		CreatedAt:   req.CreatedAt.Time,
		UpdatedAt:   req.CreatedAt.Time,
	}

	err := d.Movie.CreateMovie(c.Request().Context(), movie)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Internal server error",
			"error":   err.Error(),
		})
	}

	return c.NoContent(http.StatusNoContent)
}

func (d *Dependency) UpdateMovie(c echo.Context) error {
	movieId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "invalid movie id",
			"error":   err.Error(),
		})
	}

	if movieId == 0 {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "movie id is required",
		})
	}

	var req request
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "invalid movie data",
			"error":   err.Error(),
		})
	}

	movie := &movie.Movies{
		Title:       req.Title,
		Description: req.Description,
		Rating:      req.Rating,
		Image:       req.Image,
		CreatedAt:   req.CreatedAt.Time,
		UpdatedAt:   req.CreatedAt.Time,
	}

	err = d.Movie.UpdateMovie(c.Request().Context(), movieId, movie)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Internal server error",
			"error":   err.Error(),
		})
	}

	return c.NoContent(http.StatusNoContent)
}

func (d *Dependency) DeleteMovie(c echo.Context) error {
	movieId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "invalid movie id",
			"error":   err.Error(),
		})
	}

	if movieId == 0 {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "movie id is required",
		})
	}

	err = d.Movie.DeleteMovie(c.Request().Context(), movieId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Internal server error",
			"error":   err.Error(),
		})
	}

	return c.NoContent(http.StatusNoContent)
}
