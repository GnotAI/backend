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

    return c.Status(200).JSON(powerups)

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

    return c.Status(200).JSON(users)
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
        powerups = append(powerups[:i], powerups[i+1:]...)
        return c.Status(200).SendString("Successfully deleted the powerup")
      }
    }

    return c.Status(404).JSON(fiber.Map{"error":"Todo not found"})

  })
  
  log.Fatal(app.Listen(":"+PORT))

}
