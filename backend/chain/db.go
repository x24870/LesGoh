package chain

import (
	"fmt"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresSQL() (*gorm.DB, error) {
	dbUser := os.Getenv("POSTGRES_USER")
	dbUser = "test"
	if dbUser == "" {
		return nil, fmt.Errorf("empty env POSTGRES_USER")
	}
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbHost := os.Getenv("POSTGRES_HOST")
	dbHost = "0.0.0.0"
	if dbHost == "" {
		return nil, fmt.Errorf("empty env POSTGRES_HOST")
	}
	dbPort := os.Getenv("POSTGRES_PORT")
	dbPort = "5432"
	if dbPort == "" {
		return nil, fmt.Errorf("empty env POSTGRES_PORT")
	}
	dbName := os.Getenv("POSTGRES_DATABASE")
	dbName = "test"
	if dbName == "" {
		return nil, fmt.Errorf("empty env POSTGRES")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, dbUser, dbPassword, dbName, dbPort)
	engine, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db, err := engine.DB()
	if err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(time.Duration(10) * time.Second)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return engine, nil
}
