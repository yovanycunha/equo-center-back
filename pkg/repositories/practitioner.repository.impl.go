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
	var practitionerFound models.Practitioner

	query := bson.D{bson.E{Key: "document", Value: practitioner.Document}}
	errFound := pr.practitionerColl.FindOne(pr.ctx, query).Decode(&practitionerFound)
	
	if errFound == nil {
		return errors.New("practitioner with this document already exists")
	}

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
		return nil, errors.New("no practitioners found")
	}

	return practitioners, nil
}

func (pr *PractitionerRepositoryImpl) UpdatePractitioner(practitioner *models.Practitioner) error {
	filter := bson.D{bson.E{Key: "document", Value: practitioner.Document}}
	update := bson.D{bson.E{Key: "$set", Value: bson.D{
		bson.E{Key: "name", Value: practitioner.Name},
		bson.E{Key: "birthdate", Value: practitioner.BirthDate},
		bson.E{Key: "age", Value: practitioner.Age},
		bson.E{Key: "gender", Value: practitioner.Gender},
		bson.E{Key: "admissiondate", Value: practitioner.AdmissionDate},
		bson.E{Key: "cid", Value: practitioner.CID},
		bson.E{Key: "sponsor", Value: practitioner.Sponsor},
		bson.E{Key: "address", Value: practitioner.Address},
	}}}
	
	result, _ := pr.practitionerColl.UpdateOne(pr.ctx, filter, update)
	if result.MatchedCount != 1 {
		return errors.New("practitioner not found")
	}
	
	return nil
}

func (pr *PractitionerRepositoryImpl) DeletePractitioner(document *string) error {
	filter := bson.D{bson.E{Key: "document", Value: document}}
	result, _ := pr.practitionerColl.DeleteOne(pr.ctx, filter)
	if result.DeletedCount != 1 {
		return errors.New("practitioner not found")
	}

	return nil
}
