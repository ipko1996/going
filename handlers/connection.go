package handlers

import "gorm.io/gorm"

var DB *gorm.DB

// SetDatabase allows the main package to set the database connection
func SetDatabase(database *gorm.DB) {
	DB = database
}
