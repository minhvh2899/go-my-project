// cmd/server/main.go
package main

import (
	"my-project/internal/api"
	"my-project/internal/config"
	"my-project/internal/database"
	"my-project/pkg/logger"

	"github.com/gin-gonic/gin"
)

func main() {
	log := logger.NewLogger()

	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load configuration")
	}

	// Initialize database
	db, err := database.Initialize(cfg.DatabaseURL)
	if err != nil {
		log.Fatal().Err(err).Msgf("Failed to initialize database: %s", cfg.DatabaseURL)
	}

	// Run migrations
	if err := database.Migrate(db); err != nil {
		log.Fatal().Err(err).Msg("Failed to run database migrations")
	}

	// Initialize router
	router := gin.Default()

	// Setup API
	api.SetupRoutes(router, db)

	// Start server
	err = router.Run(cfg.ServerAddress)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to start server")
	}
}