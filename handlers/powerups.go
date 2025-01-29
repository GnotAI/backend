package handlers

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"project/backend/db"
	"project/backend/models"
)

var powerups = []models.Powerups{}

// Get all creeated powerups
func GetAllPowerups(c fiber.Ctx) error {
	// Access the MongoDB client
	client := db.Cli
	if client == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Database connection not initialized",
		})
	}

	// Access the collection in the "powerups" database
	db := client.Database("powerups")
	collection := db.Collection("powerups col")

	// Fetch all documents
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Printf("MongoDB Find error: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch powerups",
		})
	}
	defer cursor.Close(context.Background())

	// Decode results into a slice of Powerups
	if err := cursor.All(context.Background(), &powerups); err != nil {
		log.Printf("Cursor decode error: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to decode powerups",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"powerups": powerups,
	})
} 

// Create new powerup
func CreatePowerup (c fiber.Ctx) error {
    powerup := new(models.Powerups)

    if err := c.Bind().JSON(powerup); err != nil {
      return err
    }

    if powerup.Duration == 0 || powerup.Name == "" {
      return c.Status(400).JSON(fiber.Map{"error":"Powerup requires a name and duration"})
    }


  powerup.ID = primitive.NewObjectID()

  if err := addPowerup(db.Cli, "powerups", powerup); err != nil {
    return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
      "error": "Failed to create powerup",
    })
  }

  return c.Status(fiber.StatusCreated).JSON(fiber.Map{
    "message": "Powerup created successfully",
    "powerup":    powerup,
  })

}

// Creates new powerup in db
func addPowerup(client *mongo.Client, dbName string, powerup *models.Powerups) error {
	if client == nil {
		log.Println("MongoDB client is nil")
	}

	dataB := client.Database(dbName)
	collection := dataB.Collection("powerups col")
	_, err := collection.InsertOne(context.Background(), powerup)
	if err != nil {
    log.Printf("MongoDB InsertOne error: %v", err)
    return err
	}

  return nil
}

// Update powerups by ID
func UpdatePowerup (c fiber.Ctx) error {
  // Get and validate ID parameter
  id := c.Params("id")
  if id == "" {
    return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
      "error": "Powerup ID is required",
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

  // Access collection (assuming your collection is named "powerups")
  collection := client.Database("powerups").Collection("powerups col")

  // Create filter and update
  filter := bson.M{"_id": objectID}
  update := bson.M{"$set": updateData}

  // Perform update
  result, err := collection.UpdateOne(context.Background(), filter, update)
  if err != nil {
    log.Printf("MongoDB UpdateOne error: %v", err)
    return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
      "error": "Failed to update powerup",
    })
  }

  if result.MatchedCount == 0 {
    return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
      "error": "Powerup not found",
    })
  }

  return c.Status(fiber.StatusOK).JSON(fiber.Map{
    "message": "Powerup updated successfully",
  })
}

// Delete powerups by ID
func DeletePowerup(c fiber.Ctx) error {
  // Get and validate ID parameter
  id := c.Params("id")
  if id == "" {
    return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
      "error": "Powerup ID is required",
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
  collection := client.Database("powerups").Collection("powerups col")

  // Create filter
  filter := bson.M{"_id": objectID}

  // Perform deletion
  result, err := collection.DeleteOne(context.Background(), filter)
  if err != nil {
    log.Printf("MongoDB DeleteOne error: %v", err)
    return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
      "error": "Failed to delete powerup",
    })
  }

  // Check if no document was deleted
  if result.DeletedCount == 0 {
    return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
      "error": "Powerup not found",
    })
  }

  return c.Status(fiber.StatusOK).JSON(fiber.Map{
    "message": "Powerup deleted successfully",
  })
}
