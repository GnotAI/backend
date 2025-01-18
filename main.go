package main

import (
	"log"
	"os"

  "project/backend/routes"
	fiber "github.com/gofiber/fiber/v3"
	gde "github.com/joho/godotenv"
)

func main() {

  app := fiber.New()

  err := gde.Load(".env")
  if err != nil {
    log.Fatal("Failed to load .env file")
  }

  PORT := os.Getenv("PORT")

  routes.UserRoutes(app)
  routes.TaskRoutes(app)
  routes.PowerupRoutes(app)
  
  log.Fatal(app.Listen(":"+PORT))

}
