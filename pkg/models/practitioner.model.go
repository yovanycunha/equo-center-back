package models

import "time"

type Practitioner struct {
	Name string `json:"name" bson: "name"`
	BirthDate time.Time `json:"birthdate" bson: "birth_date"`
	Age int `json:"age" bson: "age"`
	Gender string `json:"gender" bson:"gender" binding:"oneof=male female not_specified"`
	AdmissionDate time.Time `json:"admissiondate" bson:"admission_date"`
	CID string `json:"cid" bson:"cid"`
}
