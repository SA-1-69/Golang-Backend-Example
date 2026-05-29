package docs

import "github.com/swaggo/swag"

// SwaggerInfo holds exported Swagger API metadata.
var SwaggerInfo = &swag.Spec{
    Version:          "1.0",
    Host:             "localhost:8080",
    BasePath:         "/api/v1",
    Schemes:          []string{"http"},
    Title:            "Golang Example Backend API",
    Description:      "A production-style REST API with Gin, GORM, PostgreSQL, and JWT authentication.",
    InfoInstanceName: "swagger",
}
