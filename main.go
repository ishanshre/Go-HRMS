package main

import (
	"log"

	"github.com/ishanshre/go-hrm/database"
	"github.com/ishanshre/go-hrm/middleware"
)

func main() {
	// calling the function to connect to database
	store, err := database.NewPostgresStore()
	if err != nil {
		log.Fatalf("Error in Connecting to database: %v", err)
	}
	// calling the init function
	if err := store.Init(); err != nil {
		log.Fatalf("error in creating a table %v", err)
	}
	server := middleware.NewApiServer(":8000", store)
	server.Run()

}
