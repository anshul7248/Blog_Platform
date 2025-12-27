package database

import (
	"fmt"
	"log"

	"blog_project/config"
	"blog_project/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Connect initializes the PostgreSQL connection
func Connect(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DBHost, cfg.DBUser, cfg.DBPass, cfg.DBName, cfg.DBPort,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	log.Println("✅ Database connected successfully")
	db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`)

	// Run migrations
	if err := db.AutoMigrate(
		&models.User{},
		&models.Post{},
		&models.Comment{},
		&models.Tag{},
		&models.Reaction{},
		// &models.Media{},
	); err != nil {
		return nil, fmt.Errorf("auto-migration failed: %w", err)
	}

	DB = db
	return db, nil
}

// CloseDB closes the database connection
func CloseDB(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		log.Printf("⚠️ Failed to close DB: %v", err)
		return
	}
	sqlDB.Close()
	log.Println("🛑 Database connection closed")
}
