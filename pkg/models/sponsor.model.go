package models

type Sponsor struct {
	Name     string `json:"name" bson: "sponsor_name"`
	Document string `json:"document" bson: "sponsor_document"`
	Phone    string `json:"phone" bson: "phone"`
	Email    string `json:"email" bson: "email"`
}