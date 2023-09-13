package models

import (
	"RS-Backend/config"
)

// MigrateDB sets up and updates database tables based on the models.
func MigrateDB() {
	config.DB.AutoMigrate(&User{})
	// If there are other models, e.g. Article, migrate like so:
	// config.DB.AutoMigrate(&Article{})
}
