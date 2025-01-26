package handlers

import (
  "fmt"
  "log"

  "project/backend/db"

  "github.com/gofiber/fiber/v3"
  "project/backend/models"
)

var users = []models.Users{}

// Get all registered users
func GetAllUsers (c fiber.Ctx) error {
    return c.Status(200).JSON(users)
}

// Create new user
func CreateUser (c fiber.Ctx) error{
  var user models.Users
  if err := c.Bind().JSON(user); err != nil {
    return err
  }

  if user.Username == "" || user.Password == "" {
    return c.Status(400).JSON(fiber.Map{"error":"User requires a username and password"})
  }  

  user.ID = len(users) + 1
  users = append(users, user)

  query := `INSERT INTO users (id ,username, pssword) VALUES ($1, $2, $3)`

  err := db.DB.QueryRow(query, user.Username, user.Password)
  if err != nil {
    log.Fatalf("Failed to insert user into database")
  }

  return c.Status(200).JSON(fiber.Map{"message" : fmt.Sprintf("User %s added successfully", users[user.ID-1].Username)})
}

// Update user by ID
func UpdateUser (c fiber.Ctx) error {
    id := c.Params("id")

    for i, user := range users {
      if fmt.Sprint(user.ID) == id {
        users[i].Status = true
        user = users[i]
        return c.Status(200).SendString(fmt.Sprintf("User %s has been updated.", users[user.ID-1].Username))
      }
    }
    
    return c.Status(404).JSON(fiber.Map{"error":"Powerup not found"})

}
//Delete user by ID
func DeleteUser (c fiber.Ctx) error {
    id := c.Params("id")

    for i, user := range users {
      if fmt.Sprint(user.ID) == id {
        val_user := users[user.ID-1].Username
        users = append(users[:i], users[i+1:]...)

        for j := i; j < len(users); j++ {
          users[j].ID--
        }

        return c.Status(200).SendString(fmt.Sprintf("Successfully deleted the user %s", val_user))
      }
    }

    return c.Status(404).JSON(fiber.Map{"error":"User not found"})

}
