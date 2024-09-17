package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Activity struct {
	ID 				primitive.ObjectID 	`json:"id" bson:"_id, omitempty"`
	Title 			string 				`json:"title" bson:"title"`
	Purpose 		string 				`json:"purpose" bson:"purpose"`
	Professionals 	[]string 			`json:"professionals" bson:"professionals"`
	Actions 		string 				`json:"actions" bson:"actions"`
	Feedback 		string 				`json:"feedback" bson:"feedback"`
}