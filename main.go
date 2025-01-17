package main

import (
	"fmt"
	"log"
	"os"

	fiber "github.com/gofiber/fiber/v3"
  gde "github.com/joho/godotenv"
)

type Powerups struct {
  ID        int     `json:"id"`
  Name      string  `json:"name"`
  Duration  int     `json:"duration"`
  Active    bool    `json:"active"`
}

type Users struct {
  ID        int        `json:"id"`
  Username  string     `json:"username"`
  Password  string     `json:"password"`
}

func main() {
  app := fiber.New()

  powerups := []Powerups{}
  users := []Users{}

  err := gde.Load(".env")
  if err != nil {
    log.Fatal("Failed to load .env file")
  }

  PORT := os.Getenv("PORT")



  // Get users
  app.Get("/users", func(c fiber.Ctx) error {
    return c.Status(200).JSON(users)
  })

  // Get powerups
  app.Get("/powerups", func(c fiber.Ctx) error {
    return c.Status(200).JSON(powerups)
  })



  // Add new powerups
  app.Post("/submit/powerups", func(c fiber.Ctx) error {
    powerup := new(Powerups)

    if err := c.Bind().JSON(powerup); err != nil {
      return err
    }

    if powerup.Duration == 0 || powerup.Name == "" {
      return c.Status(400).JSON(fiber.Map{"error":"Powerup requires a name and duration"})
    }

    powerup.ID = len(powerups) + 1
    powerups = append(powerups, *powerup)

    return c.Status(200).SendString(fmt.Sprintf("Powerup %s added successfully", powerups[powerup.ID-1].Name))

  })

  // Add new users
  app.Post("/submit/users", func(c fiber.Ctx) error{
    user := new(Users)

    if err := c.Bind().JSON(user); err != nil {
      return err
    }

    if user.Username == "" || user.Password == "" {
      return c.Status(400).JSON(fiber.Map{"error":"Users require a username and password"})
    }  

    user.ID = len(users) + 1
    users = append(users, *user)

    return c.Status(200).SendString(fmt.Sprintf("User %s added successfully", users[user.ID-1].Username))
  })



  // Update powerups (mark as completed)
  app.Patch("/patch/powerups/:id", func(c fiber.Ctx) error {
    id := c.Params("id")

    for i, powerup := range powerups {
      if fmt.Sprint(powerup.ID) == id {
        powerups[i].Active = true
        powerup = powerups[i]
        return c.Status(200).JSON(powerup)
      }
    }
    
    return c.Status(404).JSON(fiber.Map{"error":"Todo not found"})

  })

  

  // Delete powerups
  app.Delete("/del/powerups/:id", func(c fiber.Ctx) error {
    id := c.Params("id")

    for i, powerup := range powerups {
      if fmt.Sprint(powerup.ID) == id {
        val_powerup := powerups[powerup.ID-1].Name
        powerups = append(powerups[:i], powerups[i+1:]...)

        for j := i; j < len(powerups); j++ {
          powerups[j].ID--
        }

        return c.Status(200).SendString(fmt.Sprintf("Successfully deleted the powerup %s", val_powerup))
      }
    }

    return c.Status(404).JSON(fiber.Map{"error":"Todo not found"})

  })

  // Delete users
  app.Delete("/del/users/:id", func(c fiber.Ctx) error {
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

    return c.Status(404).JSON(fiber.Map{"error":"Todo not found"})

  })
  
  log.Fatal(app.Listen(":"+PORT))

}
