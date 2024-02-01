package server

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"pictiv-api/internal/model"
)

func (s *Server) handleGetIllustrator(c echo.Context) error {
	var i model.IllustratorRead
	var j model.IllustratorRead
	err := c.Bind(&i)
	fmt.Println(i.Page)
	if err != nil {
		return echo.ErrBadRequest
	} else if i == j {
		// Has nothing so give back page 1
		illustrators, err := s.db.FindManyIllustrators(model.IllustratorDTO{}, 1)
		if err != nil {
			return echo.ErrInternalServerError
		}

		return c.JSON(http.StatusOK, illustrators)
	} else if i.Page > 0 {
		// Has Page so lookup page in Many
		illustrators, err := s.db.FindManyIllustrators(model.IllustratorDTO{}, i.Page)
		if err != nil {
			return echo.ErrInternalServerError
		}

		return c.JSON(http.StatusOK, illustrators)
	} else {
		// Has ID so lookup on one
		iDTO := model.IllustratorDTO{
			ID: i.ID,
		}
		illustrator, err := s.db.FindOneIllustrator(iDTO)
		if err != nil {
			return echo.ErrInternalServerError
		}

		return c.JSON(http.StatusOK, illustrator)
	}
}

func (s *Server) handleSearchIllustrator(c echo.Context) error {
	var i model.IllustratorSearch
	var j model.IllustratorSearch
	err := c.Bind(&i)
	if err != nil {
		return echo.ErrBadRequest
	} else if i == j {
		// TODO: Add recommended algorithm

		// Has nothing so give back page 1
		//illustrator, err := s.db.FindOneIllustrator(model.IllustratorDTO{})
		//if err != nil {
		//	return echo.ErrInternalServerError
		//}
		//
		//return c.JSON(http.StatusOK, illustrator)
		return c.JSON(http.StatusOK, "nothing")
	} else {
		// Has ID so lookup on one
		iDTO := model.IllustratorDTO{
			ID:        i.ID,
			Name:      i.Name,
			TwitterID: i.TwitterID,
			PixivID:   i.PixivID,
		}
		illustrator, err := s.db.FindOneIllustrator(iDTO)
		if err != nil {
			return echo.ErrInternalServerError
		}

		return c.JSON(http.StatusOK, illustrator)
	}
}
