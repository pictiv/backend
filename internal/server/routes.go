package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	middlewares "pictiv-api/internal/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", s.statusHandler)

	e.GET("/protected", s.statusHandler, middlewares.SessionMiddleware())

	return e
}

func (s *Server) statusHandler(c echo.Context) error {
	if !s.db.Health() {
		return echo.ErrInternalServerError
	}
	return c.NoContent(http.StatusOK)
}
