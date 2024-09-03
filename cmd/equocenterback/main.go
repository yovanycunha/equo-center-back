package main

import (
	"context"
	"equocenterback/initializers"
	"equocenterback/pkg/controllers"
	"equocenterback/pkg/repositories"
	"fmt"
	"log"

	"github.com/gin-contrib/cors"
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
	PractitionerService 	repositories.PractitionerRepository
	practitionerController 	controllers.PractitionerController
)

func init() {
	initializers.LoadEnv()

	ctx := context.TODO()

	mongoConn := options.Client().ApplyURI("mongodb+srv://admin:416n3QPWlJuISzrC@equocenter.2c5e2.mongodb.net/?retryWrites=true&w=majority&appName=equocenter")
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
	PractitionerService = repositories.New(practitionerColl, ctx)
	practitionerController = controllers.New(PractitionerService)



	server = gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization", "Accept", "User-Agent", "Access-Control-Allow-Origin"}
    config.ExposeHeaders = []string{"Content-Length"}
    config.AllowCredentials = true

	server.Use(cors.New(config))

}

func main() {
	defer mongoclient.Disconnect(cxt)

	basepath := server.Group("/api")
	practitionerController.RegisterPractitionerRoutes(basepath)

	log.Fatal(server.Run())
}