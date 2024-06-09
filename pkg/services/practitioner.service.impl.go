package services

import (
	"context"
	"equocenterback/pkg/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type PractitionerServiceImpl struct {
	practitionerColl *mongo.Collection
	ctx 			context.Context
}

func New(practitionerColl *mongo.Collection, ctx context.Context) *PractitionerServiceImpl {
	return &PractitionerServiceImpl{
		practitionerColl: 	practitionerColl,
		ctx: 				ctx,
	}
}

func (ps *PractitionerServiceImpl) CreatePractitioner(practitioner *models.Practitioner) error {
	_, err := ps.practitionerColl.InsertOne(ps.ctx, practitioner)

	return err
}