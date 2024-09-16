package repositories

import (
	"context"
	"equocenterback/pkg/models"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ActivityRepositoryImpl struct {
	activityColl 	*mongo.Collection
	ctx 			context.Context
}

func NewActivityRepository(activityColl *mongo.Collection, ctx context.Context) *ActivityRepositoryImpl {
	return &ActivityRepositoryImpl{
		activityColl: activityColl,
		ctx: ctx,
	}
}

func (ar *ActivityRepositoryImpl) CreateActivity(activity *models.Activity) error {
	activity.ID = primitive.NewObjectID()
	_, err := ar.activityColl.InsertOne(ar.ctx, activity)
	return err
}

func (ar *ActivityRepositoryImpl) GetAllActivities() ([]*models.Activity, error) {
	var activities []*models.Activity

	cursor, err := ar.activityColl.Find(ar.ctx, primitive.D{})
	if err != nil {
		return nil, err
	}

	for cursor.Next(ar.ctx) {
		var activity models.Activity

		err := cursor.Decode(&activity)
		if err != nil {
			return nil, err
		}

		activities = append(activities, &activity)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(ar.ctx)

	if len(activities) == 0 {
		return nil, errors.New("no activities found")
	}

	return activities, nil
}

func (ar *ActivityRepositoryImpl) GetActivity(id *string) (*models.Activity, error) {
	var activity models.Activity
	objId, errorObj := primitive.ObjectIDFromHex(*id)
	if errorObj != nil {
		return nil, errorObj
	}

	query := bson.D{bson.E{Key: "_id", Value: objId}}
	err := ar.activityColl.FindOne(ar.ctx, query).Decode(&activity)

	return &activity, err
}

func (ar *ActivityRepositoryImpl) UpdateActivity(activity *models.Activity) error {
	objectId, errorObj := primitive.ObjectIDFromHex(activity.ID.Hex())
	if errorObj != nil {
		return errorObj
	}

	filter := bson.D{bson.E{Key: "_id", Value: objectId}}
	update := bson.D{bson.E{Key: "$set", Value: bson.D{
		bson.E{Key: "purpose", Value: activity.Purpose},
		bson.E{Key: "professionals", Value: activity.Professionals},
		bson.E{Key: "actions", Value: activity.Actions},
		bson.E{Key: "feedback", Value: activity.Feedback},
	}}}

	result,_ := ar.activityColl.UpdateOne(ar.ctx, filter, update)
	if result.MatchedCount != 1 {
		return errors.New("no activity found")
	}

	return nil
}