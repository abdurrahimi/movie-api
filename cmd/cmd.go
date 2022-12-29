package cmd

import (
	"database/sql"
	"movie-api/movie"
	moviehttp "movie-api/movie/httpserver"

	"github.com/labstack/echo/v4"
)

type dependency struct {
	movie *moviehttp.Dependency
}

func New(
	db *sql.DB,
) *dependency {
	return &dependency{
		movie: &moviehttp.Dependency{
			Movie: &movie.Dependency{
				DB: db,
			},
		},
	}
}

func (d *dependency) MovieHandler(r *echo.Group) {
	d.movie.HttpServer(r)
}
