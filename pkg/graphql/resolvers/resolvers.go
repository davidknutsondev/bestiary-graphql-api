package resolvers

import (
	// "context"
	"database/sql"
	// "errors"
	"github.com/graphql-go/graphql"
	// "github.com/davidknutsondev/bestiary-graphql-api/pkg/graphql/schema"
	// "github.com/davidknutsondev/bestiary-graphql-api/pkg/models"
	// "github.com/graphql-go/graphql"
)

// resolvers/query_resolver.go
// package resolvers

type QueryResolver struct {
	DB *sql.DB // Your database connection or any other dependencies
}

func NewQueryResolver(db *sql.DB) *QueryResolver {
	return &QueryResolver{DB: db}
}

// Define resolver methods for the root query fields here

func (r *QueryResolver) GetBeast(p graphql.ResolveParams) (interface{}, error) {
	// ctx := p.Context
	// db := r.DB

	// Your existing resolver code for "beast" field

	return nil, nil
}

func (r *QueryResolver) GetBeastList(p graphql.ResolveParams) (interface{}, error) {
	// ctx := p.Context
	// db := r.DB

	// Your existing resolver code for "beastList" field

	return nil, nil
}
