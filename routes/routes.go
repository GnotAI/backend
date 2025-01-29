package routes

import (
  "project/backend/handlers"
  fiber "github.com/gofiber/fiber/v3"
)

func UserRoutes(app *fiber.App) {
  userGroup := app.Group("/users")

  userGroup.Get("/", handlers.GetAllUsers)
  userGroup.Post("/", handlers.CreateUser)
  userGroup.Patch("/:id", handlers.UpdateUser)
  userGroup.Delete("/:id", handlers.DeleteUser)
}

func PowerupRoutes(app *fiber.App){
  powerupGroup := app.Group("/powerups")

  powerupGroup.Get("/", handlers.GetAllPowerups)
  powerupGroup.Post("/", handlers.CreatePowerup)
  powerupGroup.Patch("/:id", handlers.UpdatePowerup)
  powerupGroup.Delete("/:id", handlers.DeletePowerup)
}

func TaskRoutes(app *fiber.App){
  taskGroup := app.Group("/tasks")

  taskGroup.Get("/", handlers.GetAllTasks)
  taskGroup.Post("/", handlers.CreateTask)
  taskGroup.Patch("/:id", handlers.UpdateTask)
  taskGroup.Delete("/:id", handlers.DeleteTask)
}
