package routes

import (
	"github.com/MehulxBuilds/Go-Workflows/internal/http/handlers"
	"github.com/MehulxBuilds/Go-Workflows/internal/http/middleware"
	"github.com/gofiber/fiber/v3"
)

type RouteDependencies struct {
	AuthHandler            *handlers.AuthHandler
	AuthMiddleware         *middleware.AuthMiddleware
}

func Register(app *fiber.App, deps RouteDependencies) {

	// Create a Route Group for user authentication
	auth := app.Group("/auth")

	// /auth/...
	auth.Get("/google", deps.AuthHandler.StartGoogleAuth)
	auth.Get("/google/callback", deps.AuthHandler.GoogleCallback)
	auth.Post("/logout", deps.AuthHandler.Logout)
	auth.Get("/me", deps.AuthMiddleware.RequireAuth(), deps.AuthHandler.GetUserInfo)

	// Create a Route Group for admin only routes
	admin := app.Group("/admin", deps.AuthMiddleware.RequireAuth(), deps.AuthMiddleware.RequireAdmin())

	// /admin/...
	admin.Get("/me", deps.AuthHandler.GetUserInfo)
}