package main

import (
	"log"
	"os"

	_ "project/backend/db"
	"project/backend/routes"

	fiber "github.com/gofiber/fiber/v3"
)

func main() {

  log.Println("Starting fiber app...")
  app := fiber.New()
  PORT := os.Getenv("PORT")

  log.Println("Registering routes...")
  routes.UserRoutes(app)
  routes.TaskRoutes(app)
  routes.PowerupRoutes(app)
  
  log.Fatal(app.Listen(":"+PORT))

}
