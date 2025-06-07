package config

import (
	"log"
	"task_manager_go/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// InitDB initializes and returns a database connection.
// Sets up PostgreSQL connection with the specified configuration.
// Performs database migration for the Task model.
// Returns a GORM database instance or terminates the application on error.
func InitDB() *gorm.DB {

	connStr := "host=localhost port=5432 user=alex password=alex dbname=taskdb sslmode=disable"

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatalf("connecting is aborted %v", err)
	}

	sqlDb, err := db.DB()
	if err != nil {
		log.Fatalf("creating db is uncompleted %v", err)
	}

	err = sqlDb.Ping()
	if err != nil {
		log.Fatalf("error responce fron db %v", err)
	}
	log.Println("Database was connected")

	err = db.AutoMigrate(&model.Task{})
	if err != nil {
		log.Fatalf("Error migrating database: %v", err)
	}

	return db
}
