package models

import (
	"time"

	"gorm.io/gorm"
)

type Status string

const (
	StatusPending    Status = "pending"
	StatusInProgress Status = "in-progress"
	StatusDone       Status = "done"
)

type Task struct {
	gorm.Model
	UserID      uint       `json:"user_id" gorm:"not null"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      Status     `json:"status"`
	DueDate     *time.Time `json:"due_date"`
}

type CreateTaskInput struct {
	Title       string     `json:"title" binding:"required"`
	Description string     `json:"description"`
	DueDate     *time.Time `json:"due_date"`
}

type UpdateTaskInput struct {
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      string     `json:"status"`
	DueDate     *time.Time `json:"due_date"`
}

type PatchTaskInput struct {
	Title       *string    `json:"title"`
	Description *string    `json:"description"`
	Status      *Status    `json:"status"`
	DueDate     *time.Time `json:"due_date"`
}

func IsValidStatus(s string) bool {
	switch Status(s) {
	case StatusPending, StatusInProgress, StatusDone:
		return true
	}
	return false
}
