package main

import (
	"context"
	"log"
	"os"

	"project/backend/handlers"
	"project/backend/routes"

	fiber "github.com/gofiber/fiber/v3"
	"github.com/jackc/pgx/v5"
	gde "github.com/joho/godotenv"
)

func main() {

  app := fiber.New()

  err := gde.Load(".env")
  if err != nil {
    log.Fatal("Failed to load .env file")
  }

  var dbURL string
  ENV := os.Getenv("ENV")
  if ENV == "production" {
    dbURL = os.Getenv("RENDER_DB_URL")
  } else {
    dbURL = os.Getenv("LOCAL_DB_URL")
  }

  db, err := pgx.Connect(context.Background(), dbURL)
  if err != nil {
    log.Fatalf("Unable to connect to database: %v", err)
  }
  defer db.Close(context.Background())

  log.Println("Successfully connected to the database")

  PORT := os.Getenv("PORT")

  routes.UserRoutes(app)
  app.Get("/users", )
  routes.TaskRoutes(app)
  routes.PowerupRoutes(app)
  
  log.Fatal(app.Listen(":"+PORT))

}
