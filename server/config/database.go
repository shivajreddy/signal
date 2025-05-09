package config

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() error {
	// Create DSN string from config
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		AppConfig.DBConfig.Host,
		AppConfig.DBConfig.User,
		AppConfig.DBConfig.Password,
		AppConfig.DBConfig.Name,
		AppConfig.DBConfig.Port,
		AppConfig.DBConfig.SSLMode,
	)

	var db *gorm.DB
	var err error
	maxRetries := 5
	retryInterval := 5 * time.Second

	for i := 0; i < maxRetries; i++ {
		// Open database connection
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		}

		log.Printf("Failed to connect to database (attempt %d/%d): %v", i+1, maxRetries, err)
		if i < maxRetries-1 {
			log.Printf("Retrying in %v...", retryInterval)
			time.Sleep(retryInterval)
		}
	}

	if err != nil {
		return fmt.Errorf("failed to connect to database after all retries: %v", err)
	}

	// Get the underlying *sql.DB
	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("failed to get database instance: %v", err)
	}

	// Set connection pool settings
	sqlDB.SetMaxIdleConns(AppConfig.DBConfig.MaxIdleConns)
	sqlDB.SetMaxOpenConns(AppConfig.DBConfig.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(AppConfig.DBConfig.ConnMaxLifetime)

	DB = db
	log.Println("Database connected successfully")
	return nil
}

func GetDB() *gorm.DB {
	return DB
}
