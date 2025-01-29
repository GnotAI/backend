package db

import (
	"context"
	"log"
	"os"
  "time"

	gde "github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Cli *mongo.Client
var Collection *mongo.Collection

func Initdb() {

  log.Println("Loading .env files...")
  if err := gde.Load(".env"); err != nil {
    log.Fatal("Failed to load .env file")
  }

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
	defer client.Disconnect(ctx)

	// Create databases by inserting data
  createDatabase(client, "users")
	createDatabase(client, "powerups")
	createDatabase(client, "tasks")

  Cli = client
}

func createDatabase(client *mongo.Client, dbName string) {
	// Access the database and a collection
	db := client.Database(dbName)
	Collection := db.Collection("dummy")

	// Insert a document to persist the database
	_, err := Collection.InsertOne(context.Background(), map[string]string{"name": "init"})
	if err != nil {
		log.Fatalf("Failed to create database %s: %v", dbName, err)
	}
}
