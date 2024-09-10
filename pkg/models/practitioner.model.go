package models

import "time"

type Practitioner struct {
	Name 			string 		`json:"name" bson: "practitioner_name"`
	BirthDate 		time.Time 	`json:"birthdate" bson: "practitioner_birthdate"`
	Age 			int 		`json:"age" bson: "age"`
	Document 		string 		`json:"document" bson: "practitioner_document"`
	Gender 			string 		`json:"gender" bson:"gender"`
	AdmissionDate 	time.Time 	`json:"admissiondate" bson:"admissiondate"`
	CID 			string 		`json:"cid" bson:"cid"`

	Sponsor 		Sponsor 	`json:"sponsor" bson:"sponsor"`
	Address 		Address 	`json:"address" bson:"address"`
}

type PractitionerUpdate struct {
	Name 			string 		`json:"name" bson: "practitioner_name"`
	BirthDate 		time.Time 	`json:"birthdate" bson: "practitioner_birthdate"`
	Age 			int 		`json:"age" bson: "age"`
	NewDocument 	string 		`json:"newdocument" bson: "practitioner_newdocument"`
	Gender 			string 		`json:"gender" bson:"gender"`
	AdmissionDate 	time.Time 	`json:"admissiondate" bson:"admissiondate"`
	CID 			string 		`json:"cid" bson:"cid"`

	Sponsor 		Sponsor 	`json:"sponsor" bson:"sponsor"`
	Address 		Address 	`json:"address" bson:"address"`

	OldDocument 	string 		`json:"olddocument" bson:"olddocument"`
}