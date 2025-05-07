package main

import (
	"OneProject/db"
	"OneProject/handlers"
	"OneProject/internal/TaskService"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

func main() {

	database, err := db.InitDB()
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	e := echo.New()

	taskRepo := TaskService.NewRepository(database)
	taskServ := TaskService.NewTaskService(taskRepo)
	taskHand := handlers.NewRequestBodyHandlers(taskServ)

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/task", taskHand.GetHandler)
	e.POST("/task", taskHand.PostHandler)
	e.PATCH("/task/:id", taskHand.PatchHandler)
	e.DELETE("/task/:id", taskHand.DeleteHandler)

	e.Logger.Fatal(e.Start(":8080"))

}
