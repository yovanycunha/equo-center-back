package repositories

import (
	"context"
	"equocenterback/pkg/models"

	"go.mongodb.org/mongo-driver/bson"
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

func (pr *PractitionerRepositoryImpl) CreatePractitioner(practitioner *models.Practitioner) error {
	_, err := pr.practitionerColl.InsertOne(pr.ctx, practitioner)

	return err
}

func (pr *PractitionerRepositoryImpl) GetPractitioner(document *string) (*models.Practitioner, error) {
	var practitioner models.Practitioner

	query := bson.D{bson.E{Key: "document", Value: document}}
	err := pr.practitionerColl.FindOne(pr.ctx, query).Decode(&practitioner)

	return &practitioner, err
}

