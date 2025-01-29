package db

import (
	"context"
	"log"
	"os"
  "time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Cli *mongo.Client

func Initdb() {

  log.Println("Initializing database... ")
  connectToMongo()
  log.Println("Successfully connected to MONGOB Atlas")

}

func connectToMongo() {
  var connString string
  connString = os.Getenv("MONGODB_URL")

  // Connect to MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connString))
	if err != nil {
		log.Fatal(err)
	}

  Cli = client
}

func DisconnectMongo() {
	if Cli != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		Cli.Disconnect(ctx)
		log.Println("Disconnected from MongoDB")
	}
}
