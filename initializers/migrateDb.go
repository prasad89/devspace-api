package initializers

import (
	"log"

	"github.com/prasad89/devspace-api/models"
)

// MigrateDB applies database migrations for the DevSpace model.
func MigrateDB() {
	if err := DB.AutoMigrate(&models.DevSpace{}); err != nil {
		log.Fatal("❌ Failed to migrate database:", err)
	}
	log.Println("✅ Database migrated successfully!")
}
