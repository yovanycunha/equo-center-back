package models

type Professional struct {
	Name 			string 		`json:"name" bson: "professional_name"`
	Document 		string 		`json:"document" bson: "professional_document"`
	Specialty		string 		`json:"specialty" bson: "professional_specialty"`
}