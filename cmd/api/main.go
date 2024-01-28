package main

import (
	"fmt"
	"os"
	"pictiv-api/internal/database"
	"pictiv-api/internal/server"
)

func main() {
	if len(os.Args) < 2 {
		os.Exit(1)
	} else {
		if os.Args[1] == "serve" {
			srv := server.NewServer()

			err := srv.ListenAndServe()
			if err != nil {
				panic(fmt.Sprintf("cannot start server: %s", err))
			}
		} else if os.Args[1] == "migrate" {
			service := database.New()

			if !service.Health() {
				os.Exit(1)
			} else {
				if service.Migrate() {
					os.Exit(1)
				} else {
					fmt.Println("migrated successfully")
				}
			}
		} else {
			os.Exit(1)
		}
	}
}
