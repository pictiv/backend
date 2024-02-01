package middlewares

import (
	"context"
	"errors"
	"fmt"
	"github.com/gofrs/uuid/v5"
	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"log"
	"net/http"
	"os"
	"pictiv-api/internal/database"
	"pictiv-api/internal/model"
)

func SessionMiddleware() echo.MiddlewareFunc {
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
				log.Print(err)
				return echo.ErrUnauthorized
			}

			token, err := jwt.Parse([]byte(cookie.Value), jwt.WithKeySet(set))
			if err != nil {
				log.Print(err)
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
					i.Role = "USER"
					c.Set("user", i)
					return next(c)
				}
				return err
			} else {
				c.Set("user", user)
				return next(c)
			}
		}
	}
}
