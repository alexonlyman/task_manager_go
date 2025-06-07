package config

import (
	"context"
	"fmt"
	"task_manager_go/model"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// InitTestDBWithDocker initializes a test database using Docker container.
// Creates a PostgreSQL container for testing purposes.
// Sets up database connection and performs migrations.
// Returns a GORM database instance and a cleanup function.
func InitTestDBWithDocker() (*gorm.DB, func()) {
	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		Image:        "postgres:13",
		ExposedPorts: []string{"5432/tcp"},
		Env: map[string]string{
			"POSTGRES_USER":     "test",
			"POSTGRES_PASSWORD": "test",
			"POSTGRES_DB":       "testdb",
		},
		WaitingFor: wait.ForListeningPort("5432/tcp"),
	}
	postgresContainer, _ := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	host, _ := postgresContainer.Host(ctx)
	port, _ := postgresContainer.MappedPort(ctx, "5432/tcp")

	dsn := fmt.Sprintf("host=%s port=%s user=test password=test dbname=testdb sslmode=disable", host, port.Port())
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	err := db.AutoMigrate(&model.Task{})
	if err != nil {
		return nil, nil
	}

	cleanup := func() {
		err := postgresContainer.Terminate(ctx)
		if err != nil {
			return
		}
	}

	return db, cleanup
}
