package handlers

import (
	"context"
	"log"

	"project/backend/db"
	"project/backend/models"

	"github.com/gofiber/fiber/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var users []models.Users

// Gets all created users
func GetAllUsers(c fiber.Ctx) error {
  // Access the MongoDB client
  client := db.Cli
  if client == nil {
    return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
      "error": "Database connection not initialized",
    })
  }

  // Access the collection in the "users" database
  db := client.Database("users")
  collection := db.Collection("users col")

  // Fetch all documents
  cursor, err := collection.Find(context.Background(), bson.M{})
  if err != nil {
    log.Printf("MongoDB Find error: %v", err)
    return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
      "error": "Failed to fetch users",
    })
  }
  defer cursor.Close(context.Background())

  // Decode results into a slice of Users
  if err := cursor.All(context.Background(), &users); err != nil {
    log.Printf("Cursor decode error: %v", err)
    return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
      "error": "Failed to decode users",
    })
  }

  return c.Status(fiber.StatusOK).JSON(fiber.Map{
    "users": users,
  })
}

// Creates new user
func CreateUser(c fiber.Ctx) error {
  user := new(models.Users)

  if err := c.Bind().JSON(user); err != nil {
    return err
  }

  user.ID = primitive.NewObjectID()

  if err := addUser(db.Cli, "users", user); err != nil {
    return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
      "error": "Failed to create user",
    })
  }

  return c.Status(fiber.StatusCreated).JSON(fiber.Map{
    "message": "User created successfully",
    "user":    user,
  })
}

// Creates new user in db
func addUser(client *mongo.Client, dbName string, user *models.Users) error {
  if client == nil {
    log.Println("MongoDB client is nil")
  }

  dataB := client.Database(dbName)
  collection := dataB.Collection("users col")
  _, err := collection.InsertOne(context.Background(), user)
  if err != nil {
    log.Printf("MongoDB InsertOne error: %v", err)
    return err
  }

  return nil
}

// Updates users by ID
func UpdateUser (c fiber.Ctx) error {
  // Get and validate ID parameter
  id := c.Params("id")
  if id == "" {
    return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
      "error": "User ID is required",
    })
  }

  objectID, err := primitive.ObjectIDFromHex(id)
  if err != nil {
    return c.Status(400).JSON(fiber.Map{"error": "Invalid ID format"})
  }


  // Parse update data from request body
  var updateData map[string]interface{}
  if err := c.Bind().JSON(&updateData); err != nil {
    return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
      "error": "Invalid request body",
    })
  }

  // Access MongoDB client
  client := db.Cli
  if client == nil {
    return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
      "error": "Database connection not initialized",
    })
  }

  // Access collection (assuming your collection is named "users")
  collection := client.Database("users").Collection("users col")

  // Create filter and update
  filter := bson.M{"_id": objectID}
  update := bson.M{"$set": updateData}

  // Perform update
  result, err := collection.UpdateOne(context.Background(), filter, update)
  if err != nil {
    log.Printf("MongoDB UpdateOne error: %v", err)
    return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
      "error": "Failed to update user",
    })
  }

  if result.MatchedCount == 0 {
    return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
      "error": "User not found",
    })
  }

  return c.Status(fiber.StatusOK).JSON(fiber.Map{
    "message": "User updated successfully",
  })
}

// Delete user from db
func DeleteUser(c fiber.Ctx) error {
  // Get and validate ID parameter
  id := c.Params("id")
  if id == "" {
    return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
      "error": "User ID is required",
    })
  }

  objectID, err := primitive.ObjectIDFromHex(id)
  if err != nil {
    return c.Status(400).JSON(fiber.Map{"error": "Invalid ID format"})
  }

  // Access MongoDB client
  client := db.Cli
  if client == nil {
    return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
      "error": "Database connection not initialized",
    })
  }

  // Access collection (verify collection name matches your setup)
  collection := client.Database("users").Collection("users col")

  // Create filter
  filter := bson.M{"_id": objectID}

  // Perform deletion
  result, err := collection.DeleteOne(context.Background(), filter)
  if err != nil {
    log.Printf("MongoDB DeleteOne error: %v", err)
    return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
      "error": "Failed to delete user",
    })
  }

  // Check if no document was deleted
  if result.DeletedCount == 0 {
    return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
      "error": "User not found",
    })
  }

  return c.Status(fiber.StatusOK).JSON(fiber.Map{
    "message": "User deleted successfully",
  })
}
