// @title Golang Example Backend API
// @version 1.0
// @description REST API using Gin + GORM + PostgreSQL
// @description This API provides authentication and user management features.

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

package main

import (
    "fmt"
    "log"

    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"

    "github.com/SA/Golong-Backend-Example/internal/config"
    "github.com/SA/Golong-Backend-Example/internal/controllers"
    "github.com/SA/Golong-Backend-Example/internal/routes"
    "github.com/SA/Golong-Backend-Example/internal/utils"

    _ "github.com/SA/Golong-Backend-Example/docs"
)

func main() {
    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found, reading configuration from environment")
    }

    cfg, err := config.LoadConfig()
    if err != nil {
        log.Fatalf("failed to load configuration: %v", err)
    }

    db, err := config.ConnectDatabase(cfg)
    if err != nil {
        log.Fatalf("failed to connect to database: %v", err)
    }

    jwtProvider := utils.NewJWTProvider(cfg.JWTSecret, cfg.JWTExpiresIn)

    authController := controllers.NewAuthController(db, jwtProvider)
    userController := controllers.NewUserController(db)

    router := routes.SetupRouter(authController, userController)
    router.Use(gin.Recovery())

    port := cfg.ServerPort
    if port == "" {
        port = "8080"
    }

    log.Printf("Server is running on port %s", port)
    if err := router.Run(fmt.Sprintf(":%s", port)); err != nil {
        log.Fatalf("failed to run server: %v", err)
    }
}
