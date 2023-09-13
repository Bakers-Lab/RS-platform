package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB represents the application's connection to the database.
var DB *gorm.DB

// InitDB initializes the database connection using GORM.
func InitDB() {
	dsn := "username:password@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local" // Replace with your configuration
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database")
	}
}
