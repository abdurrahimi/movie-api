package httpserver

import (
	"movie-api/movie"

	"github.com/labstack/echo/v4"
)

type Dependency struct {
	Movie *movie.Dependency
}

func (d *Dependency) HttpServer(group *echo.Group) {
	group.GET("", d.GetMovie)
	group.GET("/:id", d.GetMovieByID)
	group.POST("", d.CreateMovie)
	group.PATCH("/:id", d.UpdateMovie)
	group.DELETE("/:id", d.DeleteMovie)
}
