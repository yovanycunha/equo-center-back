package repositories

import (
	"context"
	"equocenterback/pkg/models"
	"errors"

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