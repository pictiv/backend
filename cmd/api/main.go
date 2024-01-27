package main

import (
	"fmt"
	"pictiv-api/internal/server"
)

func main() {
	srv := server.NewServer()

	err := srv.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
