package handlers

import (
	"context"
	"log"

	"project/backend/db"
	"project/backend/models"

	"github.com/gofiber/fiber/v3"
	"go.mongodb.org/mongo-driver/mongo"
)

var users []models.Users

func CreateUser(c fiber.Ctx) error {
  user := new(models.Users)

  if err := c.Bind().JSON(user); err != nil {
    return err
  }

  if user.Username == "" || user.Password == "" {
    return c.Status(400).JSON(fiber.Map{"error":"User requires a name and password section"})
  }

  user.ID = len(users) + 1
  users = append(users, *user)

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

func addUser(client *mongo.Client, dbName string, user *models.Users) error {
	if client == nil {
		log.Println("MongoDB client is nil")
	}

	db := client.Database(dbName)
	collection := db.Collection("dummy")
	_, err := collection.InsertOne(context.Background(), *user)
	if err != nil {
    return err
	}

  return nil
}


