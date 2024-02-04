package server

import (
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	"github.com/ory/graceful"
	"net/http"
	"os"
	"pictiv-api/internal/database"
	"strconv"
)

type Server struct {
	port int
	db   database.Service
}

func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	newServer := &Server{
		port: port,
		db:   database.New(),
	}

	// Declare Server config
	server := graceful.WithDefaults(&http.Server{
		Addr:    fmt.Sprintf("127.0.0.1:%d", newServer.port),
		Handler: newServer.RegisterRoutes(),
	})

	return server
}
