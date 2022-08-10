package models

import (
	"time"
)

type Task struct {
    ID         uint      `json:"id" gorm:"primary_key"`
    AssingedTo string    `json:"assingedTo"`
    Task       string    `json:"task"`
    Deadline   string    `json:"deadline"`
    IsDone     bool      `json:"isDone"`
    CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
    UpdatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}