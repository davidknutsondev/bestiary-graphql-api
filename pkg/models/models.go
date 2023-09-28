package models

import "github.com/lib/pq"

type Beast struct {
	ID          int            `json:"id"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	OtherNames  pq.StringArray `json:"otherNames"`
	ImageURL    string         `json:"imageUrl"`
}
