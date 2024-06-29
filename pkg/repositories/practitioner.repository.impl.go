package repositories

import (
	"context"
	"equocenterback/pkg/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type PractitionerRepositoryImpl struct {
	practitionerColl 	*mongo.Collection
	ctx 				context.Context
}

func New(practitionerColl *mongo.Collection, ctx context.Context) *PractitionerRepositoryImpl {
	return &PractitionerRepositoryImpl{
		practitionerColl: 	practitionerColl,
		ctx: 				ctx,
	}
}

func (ps *PractitionerRepositoryImpl) CreatePractitioner(practitioner *models.Practitioner) error {
	_, err := ps.practitionerColl.InsertOne(ps.ctx, practitioner)

	return err
}