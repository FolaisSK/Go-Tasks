package handlers

import (
	"TaskManager/models"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	tasks  []models.Task
	nextID uint = 1
	mu     sync.Mutex
)

func GetTasks(c *gin.Context) {
	mu.Lock()
	defer mu.Unlock()
	c.JSON(http.StatusOK, gin.H{"data": tasks})
}

func CreateTask(c *gin.Context) {
	var input models.CreateTaskInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	mu.Lock()
	defer mu.Unlock()

	newTask := models.Task{
		ID:          nextID,
		Title:       input.Title,
		Description: input.Description,
		Status:      models.StatusPending,
		CreatedAt:   time.Now(),
	}
	nextID++
	tasks = append(tasks, newTask)

	c.JSON(http.StatusCreated, gin.H{"data": tasks})
}

func GetTask(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	mu.Lock()
	defer mu.Unlock()

	for _, t := range tasks {
		if t.ID == uint(id) {
			c.JSON(http.StatusOK, gin.H{"data": t})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
}

func UpdateTask(c *gin.Context) {
	var input models.UpdateTaskInput
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
	}

	mu.Lock()
	defer mu.Unlock()

	for i, t := range tasks {
		if t.ID == uint(id) {
			foundTask := models.Task{
				ID:          uint(id),
				Title:       input.Title,
				Description: input.Description,
				Status:      models.StatusPending,
			}
			tasks[i] = foundTask
			c.JSON(http.StatusOK, gin.H{"data": t})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
}

func DeleteTask(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "invalid id"})
	}
	mu.Lock()
	defer mu.Unlock()
	for i, t := range tasks {
		if t.ID == uint(id) {
			tasks = append(tasks[:i], tasks[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"data": tasks})
			return
		}
	}
	c.JSON(http.StatusNoContent, gin.H{"error": "task not found"})
}

func FilterByStatus(c *gin.Context) {
	status := c.Query("status")

	if status == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "status is required"})
		return
	}

	mu.Lock()
	defer mu.Unlock()

	filteredTasks := []models.Task{}

	for _, t := range tasks {
		if string(t.Status) == status {
			filteredTasks = append(filteredTasks, t)
		}
	}
	c.JSON(http.StatusOK, gin.H{"data": filteredTasks})
}

func PatchTask(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var input models.PatchTaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	mu.Lock()
	defer mu.Unlock()

	for i, t := range tasks {
		if t.ID == uint(id) {
			if input.Title != nil {
				tasks[i].Title = *input.Title
			}
			if input.Description != nil {
				tasks[i].Description = *input.Description
			}
			if input.Status != nil {
				tasks[i].Status = *input.Status
			}

			c.JSON(http.StatusOK, gin.H{"data": tasks[i]})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
}
