package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Activity struct {
	ID primitive.ObjectID `json:"id" bson:"_id, omitempty"`
	Purpose string `json:"purpose" bson:"activity_purpose"`
	Professionals []string `json:"professionals" bson:"activity_professionals"`
	Actions string `json:"actions" bson:"activity_actions"`
	Feedback string `json:"feedback" bson:"activity_feedback"`
}