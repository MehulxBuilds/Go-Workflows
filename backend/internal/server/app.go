package server

import (
	"github.com/MehulxBuilds/Go-Workflows/internal/config"
	"github.com/MehulxBuilds/Go-Workflows/internal/http/handlers"
	"github.com/MehulxBuilds/Go-Workflows/internal/http/middleware"
	"github.com/MehulxBuilds/Go-Workflows/internal/http/routes"
	"github.com/MehulxBuilds/Go-Workflows/internal/repositories"
	"github.com/MehulxBuilds/Go-Workflows/internal/services"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
	recovermw "github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/jackc/pgx/v5/pgxpool"
)

func New(cfg config.Config, db *pgxpool.Pool) *fiber.App {

	// New Fiber App
	app := fiber.New()

	// New Logger Mw
	app.Use(logger.New())

	// Catches if lag or failure happens
	app.Use(recovermw.New())

	// New Cors Mw
	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{cfg.FrontendURL},
		AllowCredentials: true,
		AllowMethods:     []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
	}))

	// Health Route
	app.Get("/health", health);

	// Init Repository
	userRepo := repositories.NewUserRepository(db)

	// Init Service
	authService := services.NewAuthService(cfg)

	// Init Handler
	authHandler := handlers.NewAuthHandler(cfg, authService, userRepo)

	// Init Middleware
	authMiddleware := middleware.NewAuthMiddleware(cfg, authService, userRepo)

	// Register Routes
	routes.Register(app, routes.RouteDependencies{
		AuthHandler:    authHandler,
		AuthMiddleware: authMiddleware,
	})

	return app
}
