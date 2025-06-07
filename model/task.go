package model

import "time"

// Task represents a task entity in the system.
// Used for storing task information in the database.
type Task struct {
	// Id is a unique identifier for the task
	Id uint `gorm:"primaryKey"`
	// Name is the title or name of the task
	Name string
	// Status represents the current state of the task (e.g., "Pending", "Completed")
	Status string
	// Date is the creation timestamp of the task
	Date time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}
