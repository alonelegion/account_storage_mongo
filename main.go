package main

import (
	"github.com/alonelegion/account_storage_mongo/config"
	"github.com/alonelegion/account_storage_mongo/routes"

	"github.com/gin-gonic/gin"

	"log"
)

func main() {

	// Database
	config.Connect()

	// Init Router
	router := gin.Default()

	// Route Handlers / Endpoints
	routes.Routes(router)

	log.Fatal(router.Run(":8080"))
}
