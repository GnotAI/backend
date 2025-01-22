package main

import (
	"log"
	"os"

	"project/backend/db"
	"project/backend/routes"

	fiber "github.com/gofiber/fiber/v3"
	gde "github.com/joho/godotenv"
)

func main() {

  if err := gde.Load(".env"); err != nil {
    log.Fatal("Failed to load .env file")
  }

  log.Println("Initializing database... ")
  if err := db.Initdb(); err != nil {
    log.Fatalf("Unable to initialise database: %v", err)
  }
  defer db.Pool.Close()
  log.Println("Successfully connected to the database")

  log.Println("Starting fiber app...")
  app := fiber.New()
  PORT := os.Getenv("PORT")

  log.Println("Registering routes...")
  routes.UserRoutes(app)
  routes.TaskRoutes(app)
  routes.PowerupRoutes(app)
  
  log.Fatal(app.Listen(":"+PORT))

}
