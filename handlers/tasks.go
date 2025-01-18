package handlers

import (
  "fmt"

  "github.com/gofiber/fiber/v3"
  "project/backend/models"
)

var tasks = []models.Tasks{}

// Get all tasks
func GetAllTasks (c fiber.Ctx) error {
    return c.Status(200).JSON(tasks)
}
  
// Create new tasks
func CreateTask (c fiber.Ctx) error {
    task := new(models.Tasks)

    if err := c.Bind().JSON(task); err != nil {
      return err
    }

    if task.Description == "" {
      return c.Status(400).JSON(fiber.Map{"error":"Task needs a description"})
    }

    task.ID = len(tasks) + 1
    tasks = append(tasks, *task)

    return c.Status(200).SendString(fmt.Sprintf("Task \"%s\" added successfully.", tasks[task.ID-1].Description))

}

// Update tasks by ID
func UpdateTask (c fiber.Ctx) error {
    id := c.Params("id")

    for i, task := range tasks {
      if fmt.Sprint(task.ID) == id {
        tasks[i].Completed = true
        task = tasks[i]
        return c.Status(200).SendString(fmt.Sprintf("Task \"%s\" is now completed.", tasks[task.ID-1].Description))
      }
    }
    
    return c.Status(404).JSON(fiber.Map{"error":"Task not found"})

}

// Delete tasks by ID
func DeleteTask (c fiber.Ctx) error {
    id := c.Params("id")

    for i, task := range tasks {
      if fmt.Sprint(task.ID) == id {
        task_val := tasks[task.ID-1].Description
        tasks = append(tasks[:i], tasks[i+1:]...)

        for j := i; j < len(tasks); j++ {
          tasks[j].ID--
        }

        return c.Status(200).SendString(fmt.Sprintf("Successfully deleted the task \"%s\".", task_val))
      }
    }

    return c.Status(404).JSON(fiber.Map{"error":"Powerup not found"})

  }
