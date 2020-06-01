package config

import (
	"context"
	"github.com/alonelegion/account_storage_mongo/controllers"
	"github.com/alonelegion/account_storage_mongo/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

func Connect() {

	connectionString := utils.EnvVar("DB_CONNECTION_STRING")

	// Database config
	clientOptions := options.Client().ApplyURI(connectionString)

	client, err := mongo.NewClient(clientOptions)

	// Set up a context required by mongo.Client
	ctx, cancel := context.WithTimeout(context.Background(),
		10*time.Second)

	err = client.Connect(ctx)

	// To close the connection at the end
	defer cancel()

	// Ping our DB connection
	err = client.Ping(context.Background(), readpref.Primary())

	if err != nil {
		log.Fatal("Couldn't connect to the database", err)
	} else {
		log.Println("Connected!")
	}

	// Create the database called *account_mongo_go*
	db := client.Database("account_mongo_go")

	controllers.AccountCollection(db)

	return

}
