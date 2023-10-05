package resolvers

import (
	"database/sql"
	"errors"

	"github.com/davidknutsondev/bestiary-graphql-api/pkg/models"
	"github.com/graphql-go/graphql"
)

// Define resolver functions here

// func AddBeastResolver(params graphql.ResolveParams) (interface{}, error) {
//     // Resolver logic for "addBeast" field
//     // ...
// }

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
			return nil, errors.New("Beast not found")
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
