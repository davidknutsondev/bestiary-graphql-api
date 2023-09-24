package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/davidknutsondev/bestiary-graphql-api/api"
	_ "github.com/lib/pq"
)

func main() {

	dbUser := "postgres"
	dbName := "bestiary-postgres-database"
	dbPassword := "mysecretpassword" // Use an environment variable for the password

	// Construct the connection string
	dbInfo := fmt.Sprintf("user=%s dbname=%s password=%s sslmode=disable", dbUser, dbName, dbPassword)

	// Connect to the database
	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		log.Fatal(err)
	}

	// In Go, the defer statement is used to schedule a function call to be executed just before the surrounding function returns
	// In this context it means that the db.Close() function will be called automatically when the main function exits
	// regardless of how it exits (whether normally or due to an error)
	defer db.Close()

	// Set up HTTP handlers
	http.Handle("/graphql", http.HandlerFunc(api.GraphQLHandler))
	http.Handle("/sandbox", http.HandlerFunc(api.SandboxHandler))

	// Start the HTTP server
	log.Println("Server is running on :8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
