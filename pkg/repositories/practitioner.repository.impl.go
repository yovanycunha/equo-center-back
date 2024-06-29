package repositories

import (
	"context"
	"equocenterback/pkg/models"
	"errors"

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

func (pr *PractitionerRepositoryImpl) GetAllPractitioners() ([]*models.Practitioner, error) {
	var practitioners []*models.Practitioner

	cursor, err := pr.practitionerColl.Find(pr.ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	for cursor.Next(pr.ctx) {
		 var practitioner models.Practitioner

		 err := cursor.Decode(&practitioner)
		 if err != nil {
			return nil, err
		 }

		 practitioners = append(practitioners, &practitioner)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(pr.ctx)

	if len(practitioners) == 0  {
		return nil, errors.New("No practitioners found!")
	}

	return practitioners, nil
}

