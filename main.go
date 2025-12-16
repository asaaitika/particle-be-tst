package main

import (
	"article-api/config"
	"article-api/internal/handlers"
	"article-api/internal/repository"
	"article-api/internal/validator"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	db, err := config.ConnectDatabase(cfg)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	postRepo := repository.NewPostRepository(db)
	postValidator := validator.NewValidator()
	postHandler := handlers.NewPostHandler(postRepo, postValidator)

	router := gin.Default()

	articleRoutes := router.Group("/article")
	{
		articleRoutes.POST("/", postHandler.Create)
		articleRoutes.GET("/:limit/:offset", postHandler.GetAll)
		articleRoutes.GET("/:id", postHandler.GetByID)
		articleRoutes.PUT("/:id", postHandler.Update)
		articleRoutes.DELETE("/:id", postHandler.Delete)
	}

	log.Printf("Server starting on port %s", cfg.ServerPort)
	if err := router.Run(":" + cfg.ServerPort); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
