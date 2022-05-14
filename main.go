package main

import (
	"github.com/evgreznikov/todo_app_golang/controllers"
	"github.com/evgreznikov/todo_app_golang/models"
	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()

	// Подключение к базе данных
	models.ConnectDB()

	// Маршруты
	route.GET("/tasks", controllers.GetAllTasks)
	route.POST("/tasks", controllers.CreateTask)
	route.GET("/tasks/:id", controllers.GetTaskById)
	route.PATCH("/tasks/:id", controllers.UpdateTaskById)
	route.DELETE("/tasks/:id", controllers.DeleteTaskById)

	// Запуск сервера
	route.Run()
}
