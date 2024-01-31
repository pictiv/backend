package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	middlewares "pictiv-api/internal/middleware"
	"pictiv-api/internal/model"
)

type status struct {
	token interface{}
	user  interface{}
}

func (s *Server) RegisterRoutes() http.Handler {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", s.statusHandler)

	e.GET("/protected", s.statusHandler, middlewares.SessionMiddleware())

	e.GET("/illustrators", s.illustratorQueryHandler)

	return e
}

func (s *Server) statusHandler(c echo.Context) error {
	if !s.db.Health() {
		return echo.ErrInternalServerError
	}

	resp := new(status)
	resp.user = c.Get("user")
	resp.token = c.Get("token")

	return c.JSON(http.StatusOK, resp.user)
}

func (s *Server) protectedHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, "protected")
}

func (s *Server) illustratorQueryHandler(c echo.Context) error {
	var i model.IllustratorQuery
	var j model.IllustratorQuery
	err := c.Bind(&i)
	if err != nil || i == j {
		return echo.ErrBadRequest
	}
	iDTO := model.IllustratorDTO{
		ID:        i.ID,
		Name:      i.Name,
		PixivID:   i.PixivID,
		TwitterID: i.TwitterID,
	}
	illustrator, err := s.db.FindIllustrator(iDTO)
	if err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, illustrator)
}

//func (s *Server) illustratorParamHandler(c echo.Context) error {
//	var i model.IllustratorParam
//	var j model.IllustratorParam
//	err := c.Bind(&i)
//	if err != nil || i == j {
//return echo.ErrBadRequest
//	}
//	illustrator := model.IllustratorDTO{
//		ID:        i.ID,
//		Name:      i.Name,
//		PixivID:   i.PixivID,
//		TwitterID: i.TwitterID,
//	}
//}
