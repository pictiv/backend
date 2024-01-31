package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	middlewares "pictiv-api/internal/middleware"
)

type status struct {
	token interface{}
	user  interface{}
}

func (s *Server) RegisterRoutes() http.Handler {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", s.handleStatus)

	e.GET("/protected", s.handleStatus, middlewares.SessionMiddleware())

	illustrators := e.Group("/illustrators")
	//illustrators.POST("/illustrators", )
	illustrators.GET("", s.handleGetIllustrator)
	illustrators.GET("/:id", s.handleGetIllustrator)
	//illustrators.PATCH("/illustrators/:id", )
	//illustrators.DELETE("/illustrators/:id", )
	illustrators.GET("/search", s.handleSearchIllustrator)

	return e
}

func (s *Server) handleStatus(c echo.Context) error {
	if !s.db.Health() {
		return echo.ErrInternalServerError
	}

	resp := new(status)
	resp.user = c.Get("user")
	resp.token = c.Get("token")

	return c.JSON(http.StatusOK, resp.user)
}

func (s *Server) handleProtected(c echo.Context) error {
	return c.JSON(http.StatusOK, "protected")
}
