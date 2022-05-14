package controllers

import (
	"github.com/evgreznikov/todo_app_golang/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type CreateTaskInput struct {
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description"`
	Deadline    time.Time `json:"deadline"`
}

type UpdateTaskInput struct {
	Title       string    `json:"title"`
	Completed   bool      `json:"completed"`
	Description string    `json:"description"`
	Deadline    time.Time `json:"deadline"`
}

// GetAllTasks GET /tasks
// Получаем список всех задач
func GetAllTasks(context *gin.Context) {
	var tasks []models.Task
	models.DB.Find(&tasks)

	context.JSON(http.StatusOK, gin.H{"tasks": tasks})
}

// CreateTask POST /tasks
// Добавляем новую задачу
func CreateTask(context *gin.Context) {
	var input CreateTaskInput
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task := models.Task{
		Title:       input.Title,
		Description: input.Description,
		Created_at:  time.Now(),
		Deadline:    input.Deadline,
	}
	models.DB.Create(&task)

	context.JSON(http.StatusOK, gin.H{"tasks": task})
}

// GetTaskById GET /tasks/:id
// Получаем задачу по заданному id
func GetTaskById(context *gin.Context) {
	var task models.Task
	if err := models.DB.Where("id = ?", context.Param("id")).First(&task).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Запись не существует"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"task": task})
}

// UpdateTaskById PATCH /tasks/:id
// Обновляем задачу по заданному id
func UpdateTaskById(context *gin.Context) {
	var task models.Task
	if err := models.DB.Where("id = ?", context.Param("id")).First(&task).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Запись не существует"})
		return
	}

	var input UpdateTaskInput
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&task).Update(input)

	context.JSON(http.StatusOK, gin.H{"task": task})
}

// DeleteTaskById DELETE /tasks/:id
// Удаляем задачу по заданному id
func DeleteTaskById(context *gin.Context) {
	var task models.Task
	if err := models.DB.Where("id = ?", context.Param("id")).First(&task).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Запись не существует"})
		return
	}

	var input UpdateTaskInput
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Delete(&task)

	context.JSON(http.StatusOK, gin.H{"message": "Задание успешно удалено"})
}
