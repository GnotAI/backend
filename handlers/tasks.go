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

var tasks = []models.Tasks{}

// Get all created tasks
func GetAllTasks(c fiber.Ctx) error {
	// Access the MongoDB client
	client := db.Cli
	if client == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Database connection not initialized",
		})
	}

	// Access the collection in the "tasks" database
	db := client.Database("tasks")
	collection := db.Collection("tasks col")

	// Fetch all documents
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Printf("MongoDB Find error: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch tasks",
		})
	}
	defer cursor.Close(context.Background())

	// Decode results into a slice of Tasks
	if err := cursor.All(context.Background(), &tasks); err != nil {
		log.Printf("Cursor decode error: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to decode tasks",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"tasks": tasks,
	})
}

// Create new tasks
func CreateTask (c fiber.Ctx) error {
  task := new(models.Tasks)

  if err := c.Bind().JSON(task); err != nil {
    return err
  }

  if task.Description == "" {
    return c.Status(400).JSON(fiber.Map{"error":"Task requires a description section"})
  }

  task.ID = primitive.NewObjectID()

  if err := addTask(db.Cli, "tasks", task); err != nil {
    return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
      "error": "Failed to create task",
    })
  }

  return c.Status(fiber.StatusCreated).JSON(fiber.Map{
    "message": "Task created successfully",
    "task":    task,
  })
}

// Creates new task in db
func addTask(client *mongo.Client, dbName string, task *models.Tasks) error {
	if client == nil {
		log.Println("MongoDB client is nil")
	}

	dataB := client.Database(dbName)
	collection := dataB.Collection("tasks col")
	_, err := collection.InsertOne(context.Background(), task)
	if err != nil {
    log.Printf("MongoDB InsertOne error: %v", err)
    return err
	}

  return nil
}

// Update tasks by ID
func UpdateTask (c fiber.Ctx) error {
  // Get and validate ID parameter
  id := c.Params("id")
  if id == "" {
    return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
      "error": "Task ID is required",
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

  // Access collection (assuming your collection is named "tasks")
  collection := client.Database("tasks").Collection("tasks col")

  // Create filter and update
  filter := bson.M{"_id": objectID}
  update := bson.M{"$set": updateData}

  // Perform update
  result, err := collection.UpdateOne(context.Background(), filter, update)
  if err != nil {
    log.Printf("MongoDB UpdateOne error: %v", err)
    return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
      "error": "Failed to update task",
    })
  }

  if result.MatchedCount == 0 {
    return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
      "error": "Task not found",
    })
  }

  return c.Status(fiber.StatusOK).JSON(fiber.Map{
    "message": "Task updated successfully",
  })
}

// Delete tasks by ID from db
func DeleteTask(c fiber.Ctx) error {
  // Get and validate ID parameter
  id := c.Params("id")
  if id == "" {
    return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
      "error": "Task ID is required",
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
  collection := client.Database("tasks").Collection("tasks col")

  // Create filter
  filter := bson.M{"_id": objectID}

  // Perform deletion
  result, err := collection.DeleteOne(context.Background(), filter)
  if err != nil {
    log.Printf("MongoDB DeleteOne error: %v", err)
    return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
      "error": "Failed to delete tasks",
    })
  }

  // Check if no document was deleted
  if result.DeletedCount == 0 {
    return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
      "error": "Task not found",
    })
  }

  return c.Status(fiber.StatusOK).JSON(fiber.Map{
    "message": "Task deleted successfully",
  })
}
