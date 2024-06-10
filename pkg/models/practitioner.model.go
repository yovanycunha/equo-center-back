package models

import "time"

type Practitioner struct {
	Name 			string 		`json:"name" bson: "practitioner_name"`
	BirthDate 		time.Time 	`json:"birthdate" bson: "practitioner_birthdate"`
	Age 			int 		`json:"age" bson: "age"`
	Document 		string 		`json:"document" bson: "practitioner_document"`
	Gender 			string 		`json:"gender" bson:"gender" binding:"oneof=male female not_specified"`
	AdmissionDate 	time.Time 	`json:"admissiondate" bson:"admissiondate"`
	CID 			string 		`json:"cid" bson:"cid"`

	Sponsor 		Sponsor 	`json:"sponsor" bson:"sponsor"`
	Address 		Address 	`json:"address" bson:"address"`
}
