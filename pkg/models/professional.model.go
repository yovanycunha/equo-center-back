package models

type Professional struct {
	Name 			string 		`json:"name" bson: "professional_name"`
	Document 		string 		`json:"document" bson: "professional_document"`
	Specialty		string 		`json:"specialty" bson: "professional_specialty"`
	IsCertified		bool 		`json:"iscertified" bson: "professional_is_certified"`
}

type ProfessionalUpdate struct {
	Name 			string 		`json:"name" bson: "professional_name"`
	Specialty		string 		`json:"specialty" bson: "professional_specialty"`
	IsCertified		bool 		`json:"iscertified" bson: "professional_is_certified"`
	NewDocument		string 		`json:"newdocument" bson: "professional_new_document"`
	OldDocument		string 		`json:"olddocument" bson: "professional_old_document"`
}