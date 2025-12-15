package main

import (
	"book-api-cleanarc/infrastructure/database"
	"book-api-cleanarc/infrastructure/persistence"
	httpHandlers "book-api-cleanarc/interfaces/http"
	"book-api-cleanarc/internal/command"
	"book-api-cleanarc/internal/query"

	"github.com/gin-gonic/gin"

	_ "book-api-cleanarc/docs" // مسیر docs تولید شده

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	db := database.Connect()
	database.Migrate(db)

	// Repositories
	authorRepo := persistence.NewAuthorRepo(db)
	bookRepo := persistence.NewBookRepo(db)

	// CQRS Services
	authorCmd := command.NewAuthorCommandService(authorRepo)
	authorQry := query.NewAuthorQueryService(authorRepo)
	bookCmd := command.NewBookCommandService(bookRepo)
	bookQry := query.NewBookQueryService(bookRepo)

	r := gin.Default()

	// Swagger route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Handlers
	httpHandlers.NewAuthorHandler(authorCmd, authorQry).RegisterRoutes(r)
	httpHandlers.NewBookHandler(bookCmd, bookQry).RegisterRoutes(r)

	r.Run(":8080")
}
