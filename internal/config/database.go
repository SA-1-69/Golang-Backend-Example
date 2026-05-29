package config

import (
    "fmt"

    "gorm.io/driver/postgres"
    "gorm.io/gorm"

    "github.com/SA/Golong-Backend-Example/internal/models"
)

// ConnectDatabase opens a PostgreSQL connection and runs migrations.
func ConnectDatabase(cfg *Config) (*gorm.DB, error) {
    dsn := fmt.Sprintf(
        "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
        cfg.DBHost,
        cfg.DBUser,
        cfg.DBPassword,
        cfg.DBName,
        cfg.DBPort,
    )

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        return nil, err
    }

    if err := db.AutoMigrate(&models.User{}); err != nil {
        return nil, err
    }

    return db, nil
}
