package main

import (
	"github.com/gin-gonic/gin"
	"github.com/klarqqs/dayfi_backend/internal/config"
	"github.com/klarqqs/dayfi_backend/internal/handlers"
	"github.com/klarqqs/dayfi_backend/pkg/logger"
)

func main() {
	// Initialize logger
	log := logger.New()

	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize Gin router
	r := gin.Default()

	// Routes
	r.POST("/api/payment/send", handlers.SendPayment(cfg))

	// Start server
	log.Info("Starting server on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
