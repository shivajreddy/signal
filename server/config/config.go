package config

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

// ProjectName is the name of the project
const ProjectName = "signal"

type Config struct {
	Port         string
	Env          string
	GinMode      string
	DBConfig     DBConfig
	JWTConfig    JWTConfig
	LogConfig    LogConfig
	CORSConfig   CORSConfig
	BackupConfig BackupConfig
}

type DBConfig struct {
	Host            string
	User            string
	Password        string
	Name            string
	Port            string
	SSLMode         string
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime time.Duration
}

type JWTConfig struct {
	Secret     string
	Expiration time.Duration
	Issuer     string
}

type LogConfig struct {
	Level      string
	File       string
	MaxSize    int // in MB
	MaxBackups int
	MaxAge     int // in days
}

type CORSConfig struct {
	AllowedOrigins   []string
	AllowedMethods   []string
	AllowedHeaders   []string
	AllowCredentials bool
	MaxAge           time.Duration
}

type BackupConfig struct {
	Directory     string
	RetentionDays int
	Schedule      string // cron format
	Compress      bool
}

var AppConfig Config

func LoadConfig() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	// Load server configuration
	AppConfig = Config{
		Port:    getEnv("PORT", "8080"),
		Env:     getEnv("ENVIRONMENT", "development"),
		GinMode: getEnv("GIN_MODE", "debug"),
		DBConfig: DBConfig{
			Host:            getEnv("DB_HOST", "localhost"),
			User:            getEnv("DB_USER", "postgres"),
			Password:        getEnv("DB_PASSWORD", "postgres"),
			Name:            getEnv("DB_NAME", ProjectName),
			Port:            getEnv("DB_PORT", "5432"),
			SSLMode:         getEnv("DB_SSLMODE", "disable"),
			MaxIdleConns:    getEnvInt("DB_MAX_IDLE_CONNS", 10),
			MaxOpenConns:    getEnvInt("DB_MAX_OPEN_CONNS", 100),
			ConnMaxLifetime: getEnvDuration("DB_CONN_MAX_LIFETIME", "1h"),
		},
		JWTConfig: JWTConfig{
			Secret:     getEnv("JWT_SECRET", "your-secret-key-here"),
			Expiration: getEnvDuration("JWT_EXPIRATION", "24h"),
			Issuer:     getEnv("JWT_ISSUER", ProjectName),
		},
		LogConfig: LogConfig{
			Level:      getEnv("LOG_LEVEL", "debug"),
			File:       getEnv("LOG_FILE", "./logs/app.log"),
			MaxSize:    getEnvInt("LOG_MAX_SIZE", 100),
			MaxBackups: getEnvInt("LOG_MAX_BACKUPS", 3),
			MaxAge:     getEnvInt("LOG_MAX_AGE", 30),
		},
		CORSConfig: CORSConfig{
			AllowedOrigins:   getEnvSlice("CORS_ALLOWED_ORIGINS", []string{"http://localhost:3000"}),
			AllowedMethods:   getEnvSlice("CORS_ALLOWED_METHODS", []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
			AllowedHeaders:   getEnvSlice("CORS_ALLOWED_HEADERS", []string{"Content-Type", "Authorization"}),
			AllowCredentials: getEnvBool("CORS_ALLOW_CREDENTIALS", true),
			MaxAge:           getEnvDuration("CORS_MAX_AGE", "24h"),
		},
		BackupConfig: BackupConfig{
			Directory:     getEnv("BACKUP_DIR", "./database/backups"),
			RetentionDays: getEnvInt("BACKUP_RETENTION_DAYS", 7),
			Schedule:      getEnv("BACKUP_SCHEDULE", "0 0 * * *"), // Daily at midnight
			Compress:      getEnvBool("BACKUP_COMPRESS", true),
		},
	}

	// Set Gin mode based on environment
	if AppConfig.Env == "production" {
		AppConfig.GinMode = "release"
	}
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func getEnvInt(key string, defaultValue int) int {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	var result int
	_, err := fmt.Sscanf(value, "%d", &result)
	if err != nil {
		return defaultValue
	}
	return result
}

func getEnvBool(key string, defaultValue bool) bool {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return strings.ToLower(value) == "true"
}

func getEnvDuration(key string, defaultValue string) time.Duration {
	value := os.Getenv(key)
	if value == "" {
		value = defaultValue
	}
	duration, err := time.ParseDuration(value)
	if err != nil {
		log.Printf("Error parsing duration for %s: %v, using default: %s", key, err, defaultValue)
		duration, _ = time.ParseDuration(defaultValue)
	}
	return duration
}

func getEnvSlice(key string, defaultValue []string) []string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return strings.Split(value, ",")
}

