package middleware

import (
	"strings"

	"aygit-muhasebe-integration/internal/models"
	"aygit-muhasebe-integration/pkg/errors"

	"github.com/gofiber/fiber/v2"
)

// AuthRequired is a middleware that checks for a valid session or token
func AuthRequired(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return errors.NewError(fiber.StatusUnauthorized, errors.ErrCodeUnauthorized, "Authorization header is required")
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return errors.NewError(fiber.StatusUnauthorized, errors.ErrCodeUnauthorized, "Invalid authorization format")
	}

	token := parts[1]
	// TODO: Implement actual token validation (e.g., JWT)
	if token == "" {
		return errors.NewError(fiber.StatusUnauthorized, errors.ErrCodeUnauthorized, "Invalid token")
	}

	// Placeholder user for demonstration
	// In a real app, this would be fetched from the database or token claims
	user := &models.User{
		Role: models.RoleAdmin,
	}
	c.Locals("user", user)

	return c.Next()
}

// RoleAllowed restricts access to specific roles
func RoleAllowed(roles ...models.UserRole) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user, ok := c.Locals("user").(*models.User)
		if !ok {
			return errors.NewError(fiber.StatusUnauthorized, errors.ErrCodeUnauthorized, "User context not found")
		}

		for _, role := range roles {
			if user.Role == role {
				return c.Next()
			}
		}

		return errors.NewError(fiber.StatusForbidden, "AUTH_002", "Access denied for this role")
	}
}
