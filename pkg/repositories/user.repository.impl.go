package repositories

import (
	"context"
	"equocenterback/pkg/models"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepositoryImpl struct {
	userColl *mongo.Collection
	ctx      context.Context
}

func NewUserRepo(userColl *mongo.Collection, ctx context.Context) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		userColl: userColl,
		ctx:      ctx,
	}
}

func (ur *UserRepositoryImpl) CreateUser(user *models.User) error {
	var userFound models.User

	query := bson.D{bson.E{Key: "email", Value: user.Email}}
	errFound := ur.userColl.FindOne(ur.ctx, query).Decode(&userFound)

	if errFound == nil {
		return errors.New("user with this email already exists")
	}

	user.ID = primitive.NewObjectID()
	_, err := ur.userColl.InsertOne(ur.ctx, user)
	
	return err
}

func (ur *UserRepositoryImpl) GetUser(email *string) (*models.User, error) {
	var user models.User

	query := bson.D{bson.E{Key: "email", Value: email}}
	err := ur.userColl.FindOne(ur.ctx, query).Decode(&user)

	return &user, err
}