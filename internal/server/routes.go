package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	middlewares "pictiv-api/internal/middleware"
	"pictiv-api/internal/model"
)

func (s *Server) RegisterRoutes() http.Handler {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", s.handleStatus)

	e.GET("/protected", s.handleStatus, middlewares.SessionMiddleware())
	e.GET("/admin", s.handleStatus, middlewares.SessionMiddleware())

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

	resp := model.Status{
		Token: c.Get("token").(string),
		User:  c.Get("user").(model.UserDTO),
	}

	return c.JSON(http.StatusOK, resp.User)
}
