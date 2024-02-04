package mw

import (
	"context"
	"errors"
	"fmt"
	"github.com/gofrs/uuid/v5"
	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"net/http"
	"os"
	"pictiv-api/internal/database"
	"pictiv-api/internal/model"
)

func AuthMiddleware(ROLE model.Role) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cookie, err := c.Cookie("hanko")
			if errors.Is(err, http.ErrNoCookie) {
				return echo.ErrUnauthorized
			}
			if err != nil {
				return err
			}
			set, err := jwk.Fetch(
				context.Background(),
				fmt.Sprintf("%v/.well-known/jwks.json", os.Getenv("HANKO_API")),
			)
			if err != nil {
				return echo.ErrUnauthorized
			}

			token, err := jwt.Parse([]byte(cookie.Value), jwt.WithKeySet(set))
			if err != nil {
				return echo.ErrUnauthorized
			}

			c.Set("token", cookie.Value)
			//c.Set("user", token.Subject())

			db := database.New()
			defer db.Close()

			i := model.UserDTO{ID: uuid.FromStringOrNil(token.Subject())}
			user, err := db.FindOneUser(i)
			if err != nil {
				if errors.Is(err, pgx.ErrNoRows) {
					err := db.CreateUser(i)
					if err != nil {
						return err
					}
					if model.Roles[i.Role] <= model.Roles[ROLE] {
						i.Role = model.USER
						c.Set("user", i)
						return next(c)
					} else {
						return echo.ErrForbidden
					}
				}
				return err
			} else {
				if model.Roles[user.Role] <= model.Roles[ROLE] {
					c.Set("user", user)
					return next(c)
				} else {
					return echo.ErrForbidden
				}
			}
		}
	}
}
