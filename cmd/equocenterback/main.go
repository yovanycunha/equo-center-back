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
	
	professionalColl 		*mongo.Collection
	professionalService		repositories.ProfessionalRepository
	professionalController 	controllers.ProfessionalController

	activityColl 		*mongo.Collection
	activityService		repositories.ActivityRepository
	activityController 	controllers.ActivityController

	userColl 		*mongo.Collection
	userService		repositories.UserRepository
	userController 	controllers.UserController
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

	professionalColl = mongoclient.Database("equocenter").Collection("professional")
	professionalService = repositories.NewProfessionalRepo(professionalColl, ctx)
	professionalController = controllers.NewProfessionalController(professionalService)

	activityColl = mongoclient.Database("equocenter").Collection("activity")
	activityService = repositories.NewActivityRepository(activityColl, ctx)
	activityController = controllers.NewActivityController(activityService)

	userColl = mongoclient.Database("equocenter").Collection("user")
	userService = repositories.NewUserRepo(userColl, ctx)
	userController = controllers.NewUserController(userService)

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
	professionalController.RegisterProfessionalRoutes(basepath)
	activityController.RegisterActivityRoutes(basepath)
	userController.RegisterUserRoutes(basepath)

	log.Fatal(server.Run())
}