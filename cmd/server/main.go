package main

import (
    "fmt"
    "log"

    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"

    _ "github.com/SA/Golang-Backend-Example/docs"
    "github.com/SA/Golang-Backend-Example/internal/config"
    "github.com/SA/Golang-Backend-Example/internal/controllers"
    "github.com/SA/Golang-Backend-Example/internal/routes"
    "github.com/SA/Golang-Backend-Example/internal/utils"
)

// @title Golang Backend API
// @version 1.0
// @description A simple backend API written in Go using Gin framework
// @host localhost:8080
// @BasePath /api/v1
// @schemes http https
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Enter a valid jwt token to proceed
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
