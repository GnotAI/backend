package handlers

import (
  "fmt"

  "github.com/gofiber/fiber/v3"
  "project/backend/models"
)

var powerups = []models.Powerups{}

// Get all powerups
func GetAllPowerups (c fiber.Ctx) error {
    return c.Status(200).JSON(powerups)
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

    powerup.ID = len(powerups) + 1
    powerups = append(powerups, *powerup)

    return c.Status(200).SendString(fmt.Sprintf("Powerup %s added successfully.", powerups[powerup.ID-1].Name))

}

// Update powerups by ID
func UpdatePowerup (c fiber.Ctx) error {
    id := c.Params("id")

    for i, powerup := range powerups {
      if fmt.Sprint(powerup.ID) == id {
        powerups[i].Active = true
        powerup = powerups[i]
        return c.Status(200).SendString(fmt.Sprintf("Powerup %s is now active.", powerups[powerup.ID-1].Name))
      }
    }
    
    return c.Status(404).JSON(fiber.Map{"error":"Powerup not found"})

}

// Delete powerups by ID
func DeletePowerup (c fiber.Ctx) error {
    id := c.Params("id")

    for i, powerup := range powerups {
      if fmt.Sprint(powerup.ID) == id {
        val_powerup := powerups[powerup.ID-1].Name
        powerups = append(powerups[:i], powerups[i+1:]...)

        for j := i; j < len(powerups); j++ {
          powerups[j].ID--
        }

        return c.Status(200).SendString(fmt.Sprintf("Successfully deleted the powerup %s.", val_powerup))
      }
    }

    return c.Status(404).JSON(fiber.Map{"error":"Powerup not found"})

  }
