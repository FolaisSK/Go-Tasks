package handlers

import (
	"TaskManager/db"
	"TaskManager/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTasks(c *gin.Context) {
	user, _ := c.MustGet("currentUser").(models.User)

	var tasks []models.Task
	query := db.DB.Model(&models.Task{}).Where("user_id = ?", user.ID)
	if status := c.Query("status"); status != "" {
		query = query.Where("status = ?", status)
	}

	if err := query.Find(&tasks).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": tasks})
}

func CreateTask(c *gin.Context) {
	var input models.CreateTaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, _ := c.MustGet("currentUser").(models.User)
	task := models.Task{
		Title:       input.Title,
		Description: input.Description,
		Status:      models.StatusPending,
		DueDate:     input.DueDate,
		UserID:      user.ID,
	}
	if err := db.DB.Create(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": task})
}

func GetTask(c *gin.Context) {
	user, _ := c.MustGet("currentUser").(models.User)

	var task models.Task
	if err := db.DB.First(&task, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
		return
	}

	if task.UserID != user.ID {
		c.JSON(http.StatusForbidden, gin.H{"error": "not your task"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": task})
}

func UpdateTask(c *gin.Context) {
	user, _ := c.MustGet("currentUser").(models.User)

	var task models.Task
	if err := db.DB.First(&task, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
		return
	}

	if task.UserID != user.ID {
		c.JSON(http.StatusForbidden, gin.H{"error": "not your task"})
		return
	}

	var input models.UpdateTaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.Status != "" && !models.IsValidStatus(input.Status) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": fmt.Sprintf("invalid status %q: must be one of pending, in-progress, done", input.Status),
		})
		return
	}

	if err := db.DB.Model(&task).Updates(input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": task})
}

func DeleteTask(c *gin.Context) {
	user, _ := c.MustGet("currentUser").(models.User)

	var task models.Task
	if err := db.DB.First(&task, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
		return
	}

	if task.UserID != user.ID {
		c.JSON(http.StatusForbidden, gin.H{"error": "not your task"})
		return
	}

	if err := db.DB.Delete(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

func FilterByStatus(c *gin.Context) {
	user, _ := c.MustGet("currentUser").(models.User)

	status := c.Query("status")
	if status == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "status is required"})
		return
	}

	var tasks []models.Task
	if err := db.DB.Where("user_id = ? AND status = ?", user.ID, status).Find(&tasks).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": tasks})
}

func PatchTask(c *gin.Context) {
	user, _ := c.MustGet("currentUser").(models.User)

	var task models.Task
	if err := db.DB.First(&task, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
		return
	}

	if task.UserID != user.ID {
		c.JSON(http.StatusForbidden, gin.H{"error": "not your task"})
		return
	}

	var input models.PatchTaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.DB.Model(&task).Updates(input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": task})
}
