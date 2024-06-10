package models

type Address struct {
	City    	string `json:"city" bson: "city"`
	Street 		string `json:"street" bson: "street"`
	Neiborhood 	string `json:"neiborhood" bson: "neiborhood"`
}