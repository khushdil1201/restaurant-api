package migrations

import (
    "log"
    "restaurant-api/models"
    "gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
    err := db.AutoMigrate(&models.User{}, &models.MenuItem{})
    if err != nil {
        log.Fatalf("Failed to run migrations: %v", err)
    }
}
