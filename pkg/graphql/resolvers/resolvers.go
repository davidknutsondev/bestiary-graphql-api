package resolvers

import (
	"database/sql"

	"github.com/davidknutsondev/bestiary-graphql-api/pkg/graphql/schema"
	"github.com/davidknutsondev/bestiary-graphql-api/pkg/models"
	"github.com/graphql-go/graphql"
)

type Resolver struct {
	DB *sql.DB
}

func ResolveBeast(params graphql.ResolveParams) (interface{}, error) {
	nameQuery, isOK := params.Args["name"].(string)
	if isOK {
		// Search for el with name
		for _, beast := range schema.BeastList {
			if beast.Name == nameQuery {
				return beast, nil
			}
		}
	}

	return models.Beast{}, nil
}

func ResolveBeastList(p graphql.ResolveParams) (interface{}, error) {
	return schema.BeastList, nil
}
