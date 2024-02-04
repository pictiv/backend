package main

import (
	"fmt"
	"github.com/ory/graceful"
	"log"
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

			log.Println("main: Starting the server")
			if err := graceful.Graceful(srv.ListenAndServe, srv.Shutdown); err != nil {
				log.Fatalln("main: Failed to gracefully shutdown")
			}
			log.Println("main: Server was shutdown gracefully")
		} else if os.Args[1] == "migrate" {
			service := database.New()

			if !service.Health() {
				panic("cannot connect to database")
			} else {
				if !service.Migrate() {
					panic("migration failed")
				} else {
					fmt.Println("migrated successfully")
				}
			}
		} else {
			os.Exit(1)
		}
	}
}
