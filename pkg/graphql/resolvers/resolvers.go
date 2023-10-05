package resolvers

import (
	"database/sql"
	"errors"

	"github.com/davidknutsondev/bestiary-graphql-api/pkg/models"
	"github.com/graphql-go/graphql"
	"github.com/lib/pq"
)

func AddBeastResolver(params graphql.ResolveParams) (interface{}, error) {
	// Retrieve the database connection from the context
	db := params.Context.Value("db").(*sql.DB)

	// Extract the input arguments from ResolveParams
	name, _ := params.Args["name"].(string)
	description, _ := params.Args["description"].(string)
	imageUrl, _ := params.Args["imageUrl"].(string)

	// Convert "otherNames" argument to pq.StringArray
	otherNamesArg := params.Args["otherNames"]
	otherNamesList, ok := otherNamesArg.([]interface{})
	if !ok {
		return nil, errors.New("otherNames argument must be a list")
	}

	// Convert the list of interface to a list of strings
	var otherNames pq.StringArray
	for _, item := range otherNamesList {
		if str, ok := item.(string); ok {
			otherNames = append(otherNames, str)
		} else {
			return nil, errors.New("otherNames list must contain only strings")
		}
	}

	// Perform the database insertion
	query := `
        INSERT INTO beasts (name, description, othernames, imageurl)
        VALUES ($1, $2, $3, $4)
        RETURNING id
    `
	var id int
	err := db.QueryRow(query, name, description, pq.Array(otherNames), imageUrl).Scan(&id)
	if err != nil {
		return nil, err
	}

	// Create a new Beast instance with the generated ID
	newBeast := models.Beast{
		ID:          id,
		Name:        name,
		Description: description,
		OtherNames:  otherNames,
		ImageURL:    imageUrl,
	}

	return newBeast, nil
}

// func UpdateBeastResolver(params graphql.ResolveParams) (interface{}, error) {
//     // Resolver logic for "updateBeast" field
//     // ...
// }

func GetBeastResolver(p graphql.ResolveParams) (interface{}, error) {
	ctx := p.Context
	db := ctx.Value("db").(*sql.DB)

	// Check if an "id" argument is provided in the GraphQL query.
	id, ok := p.Args["id"].(int) // Assuming the ID is of type int

	if !ok {
		return nil, errors.New("ID argument is missing or invalid")
	}

	// Query the database for a specific beast by its ID.
	query := "SELECT id, name, description, otherNames, imageURL FROM beasts WHERE id = $1"
	row := db.QueryRow(query, id)

	var beast models.Beast
	err := row.Scan(
		&beast.ID,
		&beast.Name,
		&beast.Description,
		&beast.OtherNames,
		&beast.ImageURL,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("beast not found")
		}
		return nil, err
	}

	return beast, nil
}

func GetBeastListResolver(p graphql.ResolveParams) (interface{}, error) {

	ctx := p.Context
	db := ctx.Value("db").(*sql.DB)

	rows, err := db.Query("SELECT id, name, description, otherNames, imageURL FROM beasts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var beasts []models.Beast
	for rows.Next() {
		var beast models.Beast
		err := rows.Scan(
			&beast.ID,
			&beast.Name,
			&beast.Description,
			&beast.OtherNames,
			&beast.ImageURL,
		)
		if err != nil {
			return nil, err
		}
		beasts = append(beasts, beast)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return beasts, nil
}
