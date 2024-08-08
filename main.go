package main

import (
	"asdf/handlers"
	"asdf/models"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func main() {
	// PostgreSQL connection parameters
	// TODO: put these into .env
	host := "localhost"
	user := "myuser"
	password := "mysecretpassword"
	dbname := "mydatabase"
	port := "5432"

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,          // Don't include params in the SQL log
			Colorful:                  false,         // Disable color
		},
	)

	// Build the connection string
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)

	// Connect to the PostgreSQL database
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	// Auto migrate the Product model
	DB.AutoMigrate(&models.Product{}, &models.Category{})

	// Set the database connection for handlers
	handlers.SetDatabase(DB)

	// Set up Gin routes
	r := gin.Default()
	r.POST("/products", handlers.CreateProduct)
	r.GET("/products", handlers.GetProducts)
	r.GET("/products/:id", handlers.GetProduct)
	r.PUT("/products/:id", handlers.UpdateProduct)
	r.DELETE("/products/:id", handlers.DeleteProduct)

	r.POST("/categories", handlers.CreateCategory)
	r.GET("/categories", handlers.GetCategories)
	r.GET("/categories/:id", handlers.GetCategory)
	r.PUT("/categories/:id", handlers.UpdateCategory)
	r.DELETE("/categories/:id", handlers.DeleteCategory)

	r.Run(":8080")
}
