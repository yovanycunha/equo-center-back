package repositories

import (
	"context"
	"equocenterback/pkg/models"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProfessionalRepositoryImpl struct {
	professionalColl 	*mongo.Collection
	ctx 				context.Context
}

func NewProfessionalRepo(professionalColl *mongo.Collection, ctx context.Context) *ProfessionalRepositoryImpl {
	return &ProfessionalRepositoryImpl{
		professionalColl: professionalColl,
		ctx: 				ctx,
	}
}

func (pr *ProfessionalRepositoryImpl) CreateProfessional(professional *models.Professional) error {
	var professionalFound models.Professional

	query := bson.D{bson.E{Key: "document", Value: professional.Document}}
	errFound := pr.professionalColl.FindOne(pr.ctx, query).Decode(&professionalFound)

	if errFound == nil {
		return errors.New("professional with this document already exists")
	}

	_, err := pr.professionalColl.InsertOne(pr.ctx, professional)

	return err
}

func (pr *ProfessionalRepositoryImpl) GetProfessional(document *string) (*models.Professional, error) {
	var professional models.Professional

	query := bson.D{bson.E{Key: "document", Value: document}}
	err := pr.professionalColl.FindOne(pr.ctx, query).Decode(&professional)

	return &professional, err
}

func (pc *ProfessionalRepositoryImpl) GetAllProfessionals() ([]*models.Professional, error) {
	var professionals []*models.Professional

	cursor, err := pc.professionalColl.Find(pc.ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	for cursor.Next(pc.ctx) {
		var professional models.Professional

		err := cursor.Decode(&professional)
		if err != nil {
			return nil, err
		}

		professionals = append(professionals, &professional)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(pc.ctx)

	if len(professionals) == 0 {
		return nil, errors.New("no professionals found")
	}

	return professionals, nil
}

func (pr *ProfessionalRepositoryImpl) UpdateProfessional(professional *models.ProfessionalUpdate) error {
	filter := bson.D{bson.E{Key: "document", Value: professional.OldDocument}}
	update := bson.D{bson.E{Key: "$set", Value: bson.D{
		bson.E{Key: "name", Value: professional.Name},
		bson.E{Key: "specialty", Value: professional.Specialty},
		bson.E{Key: "document", Value: professional.NewDocument},
	}}}

	result,_ := pr.professionalColl.UpdateOne(pr.ctx, filter, update)
	if result.MatchedCount != 1 {
		return errors.New("professional not found")
	}

	return nil
}