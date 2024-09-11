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