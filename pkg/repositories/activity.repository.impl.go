package repositories

import (
	"context"
	"equocenterback/pkg/models"

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