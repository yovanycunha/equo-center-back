package main

import (
	"context"
	"equocenterback/initializers"
	"equocenterback/pkg/controllers"
	"equocenterback/pkg/services"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	server *gin.Engine
	cxt context.Context

	mongoclient *mongo.Client
	
	practitionerColl 		*mongo.Collection
	PractitionerService 	services.PractitionerService
	practitionerController 	controllers.PractitionerController
)

func init() {
	initializers.LoadEnv()

	ctx := context.TODO()

	mongoConn := options.Client().ApplyURI("mongodb://localhost:27017/")
	mongoclient, err := mongo.Connect(ctx, mongoConn)
	if err != nil {
		log.Fatal(err)
	}

	err = mongoclient.Ping(ctx,readpref.Primary())
	if err != nil{
		log.Fatal(err)
	}

	fmt.Println("MongoDB connection successfully established!")

	practitionerColl = mongoclient.Database("equocenter").Collection("practitioner")
	PractitionerService = services.New(practitionerColl, ctx)
	practitionerController = controllers.New(PractitionerService)

	server = gin.Default()
}

func main() {
	defer mongoclient.Disconnect(cxt)

	basepath := server.Group("/api")
	practitionerController.RegisterPractitionerRoutes(basepath)

	log.Fatal(server.Run())
}