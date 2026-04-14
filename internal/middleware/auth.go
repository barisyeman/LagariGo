package middleware

import (
	"github.com/barisyeman/LagariGo/internal/auth"
	"github.com/gofiber/fiber/v2"
)

// RequireAuth blocks unauthenticated requests and redirects to /login.
func RequireAuth(c *fiber.Ctx) error {
	u := auth.CurrentUser(c)
	if u == nil {
		return c.Redirect("/login")
	}
	c.Locals("user", u)
	return c.Next()
}

// RequireAdmin requires the current user to have admin role.
func RequireAdmin(c *fiber.Ctx) error {
	u := auth.CurrentUser(c)
	if u == nil {
		return c.Redirect("/login")
	}
	if !u.IsAdmin() {
		return c.Status(fiber.StatusForbidden).SendString("Forbidden")
	}
	c.Locals("user", u)
	return c.Next()
}

// AttachUser populates c.Locals("user") for every request (may be nil).
func AttachUser(c *fiber.Ctx) error {
	c.Locals("user", auth.CurrentUser(c))
	return c.Next()
}
