package schema

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/davidknutsondev/bestiary-graphql-api/pkg/graphql/resolvers"
	"github.com/davidknutsondev/bestiary-graphql-api/pkg/models"
	"github.com/graphql-go/graphql"
)

// Helper function to import json from file to map
func importJSONDataFromFile(fileName string, result interface{}) (isOK bool) {
	isOK = true
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Print("Error:", err)
		isOK = false
	}
	err = json.Unmarshal(content, result)
	if err != nil {
		isOK = false
		fmt.Print("Error:", err)
	}
	return
}

var BeastList []models.Beast
var _ = importJSONDataFromFile("./beastData.json", &BeastList)

// define custom GraphQL ObjectType `beastType` for our Golang struct `Beast`
// Note that
// - the fields in our todoType maps with the json tags for the fields in our struct
// - the field type matches the field type in our struct
var beastType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Beast",
	Fields: graphql.Fields{
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"description": &graphql.Field{
			Type: graphql.String,
		},
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"otherNames": &graphql.Field{
			Type: graphql.NewList(graphql.String),
		},
		"imageUrl": &graphql.Field{
			Type: graphql.String,
		},
	},
})

// root mutation
var rootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		"addBeast": &graphql.Field{
			Type:        beastType, // the return type for this field
			Description: "add a new beast",
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"description": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"otherNames": &graphql.ArgumentConfig{
					Type: graphql.NewList(graphql.String),
				},
				"imageUrl": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: resolvers.AddBeastResolver,
		},
		"updateBeast": &graphql.Field{
			Type:        beastType, // the return type for this field
			Description: "Update existing beast",
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"description": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
				"otherNames": &graphql.ArgumentConfig{
					Type: graphql.NewList(graphql.String),
				},
				"imageUrl": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				id, _ := params.Args["id"].(int)
				affectedBeast := models.Beast{}

				// Search list for beast with id
				for i := 0; i < len(BeastList); i++ {
					if BeastList[i].ID == id {
						if _, ok := params.Args["description"]; ok {
							BeastList[i].Description = params.Args["description"].(string)
						}
						if _, ok := params.Args["name"]; ok {
							BeastList[i].Name = params.Args["name"].(string)
						}
						if _, ok := params.Args["imageUrl"]; ok {
							BeastList[i].ImageURL = params.Args["imageUrl"].(string)
						}
						if _, ok := params.Args["otherNames"]; ok {
							BeastList[i].OtherNames = params.Args["otherNames"].([]string)
						}
						// Assign updated beast so we can return it
						affectedBeast = BeastList[i]
						break
					}
				}
				// Return affected beast
				return affectedBeast, nil
			},
		},
	},
})

// root query
// test with Sandbox at localhost:8080/sandbox
var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"beast": &graphql.Field{
			Type:        beastType,
			Description: "Get single beast",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: resolvers.GetBeastResolver,
		},

		"beastList": &graphql.Field{
			Type:        graphql.NewList(beastType),
			Description: "List of beasts",
			Resolve:     resolvers.GetBeastListResolver,
		},
	},
})

// define schema, with our rootQuery and rootMutation
var BeastSchema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    rootQuery,
	Mutation: rootMutation,
})
