package main

import (
	"context"
	"equocenterback/initializers"
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

	server = gin.Default()
}

func main() {
	log.Fatal(server.Run())
}