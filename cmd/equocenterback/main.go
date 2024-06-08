package main

import (
	"equocenterback/initializers"
	"log"

	"github.com/gin-gonic/gin"
)

var (
	server *gin.Engine
)

func init() {
	initializers.LoadEnv()

	server = gin.Default()
}

func main() {
	log.Fatal(server.Run())
}