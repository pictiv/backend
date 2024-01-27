package middlewares

import (
	"context"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"log"
	"net/http"
	"os"
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

			log.Printf("session for user '%s' verified successfully", token.Subject())

			c.Set("token", cookie.Value)
			c.Set("user", token.Subject())

			return next(c)
		}
	}
}
