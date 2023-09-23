package main

import (
	"net/http"

	"github.com/davidknutsondev/bestiary-graphql-api/api"
)

func main() {
	http.Handle("/graphql", http.HandlerFunc(api.GraphQLHandler))
	http.Handle("/sandbox", http.HandlerFunc(api.SandboxHandler))

	http.ListenAndServe(":8080", nil)
}

// func main() {
// 	h := handler.New(&handler.Config{
// 		Schema:   &schema.BeastSchema,
// 		Pretty:   true,
// 		GraphiQL: false,
// 	})

// 	http.Handle("/graphql", h)

// 	http.Handle("/sandbox", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		w.Write(sandboxHTML)
// 	}))

// 	http.ListenAndServe(":8080", nil)

// }
