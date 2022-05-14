package models

import "time"

type Task struct {
	Id          uint      `json:"id" gorm:"primary_key"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	Created_at  time.Time `json:"created_at"`
	Deadline    time.Time `json:"deadline"`
}
