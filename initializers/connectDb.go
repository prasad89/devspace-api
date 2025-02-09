package initializer

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// ConnectDB initializes and returns a PostgreSQL database connection
func ConnectDB() (*gorm.DB, error) {
	host := GetEnv("DB_HOST", "localhost")
	port := GetEnv("DB_PORT", "5432")
	user := GetEnv("DB_USER", "postgres")
	password := GetEnv("DB_PASSWORD", "password")
	dbname := GetEnv("DB_NAME", "postgres")
	sslMode := GetEnv("DB_SSLMODE", "disable")

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("❌ Failed to connect to PostgreSQL:", err)
	}

	log.Println("✅ Connected to PostgreSQL successfully!")
	return db, nil
}
