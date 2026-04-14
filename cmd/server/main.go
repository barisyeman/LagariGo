package main

import (
	"log"

	"github.com/barisyeman/LagariGo/internal/auth"
	"github.com/barisyeman/LagariGo/internal/config"
	"github.com/barisyeman/LagariGo/internal/database"
	"github.com/barisyeman/LagariGo/internal/handler"
	"github.com/barisyeman/LagariGo/internal/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	fiberlogger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	cfg := config.Load()

	if err := database.Connect(cfg); err != nil {
		log.Fatalf("db connect: %v", err)
	}
	if err := database.Migrate(); err != nil {
		log.Fatalf("db migrate: %v", err)
	}
	if err := database.Seed(cfg); err != nil {
		log.Fatalf("db seed: %v", err)
	}

	auth.Init()

	app := fiber.New(fiber.Config{
		AppName: cfg.AppName,
	})

	app.Use(recover.New())
	app.Use(fiberlogger.New())
	app.Static("/assets", "./public/assets")

	// CSRF: applies to all state-changing methods. Token surfaces via c.Locals("csrf").
	app.Use(csrf.New(csrf.Config{
		KeyLookup:      "form:_csrf",
		CookieName:     "lagarigo_csrf",
		CookieSameSite: "Lax",
		CookieHTTPOnly: true,
		ContextKey:     "csrf",
	}))

	app.Use(middleware.AttachUser)

	// Static / known routes (must register before /:slug catch-all)
	app.Get("/", handler.Home)
	app.Get("/about-us", handler.About)
	app.Get("/contact", handler.Contact)

	// Auth
	app.Get("/login", handler.ShowLogin)
	app.Post("/login", handler.DoLogin)
	app.Get("/register", handler.ShowRegister)
	app.Post("/register", handler.DoRegister)
	app.Post("/logout", handler.DoLogout)

	// Admin (protected)
	adminGroup := app.Group("/admin", middleware.RequireAdmin)
	adminGroup.Get("/", handler.AdminDashboard)

	adminGroup.Get("/pages", handler.AdminPagesIndex)
	adminGroup.Get("/pages/new", handler.AdminPagesNew)
	adminGroup.Post("/pages", handler.AdminPagesCreate)
	adminGroup.Get("/pages/:id/edit", handler.AdminPagesEdit)
	adminGroup.Post("/pages/:id", handler.AdminPagesUpdate)
	adminGroup.Post("/pages/:id/delete", handler.AdminPagesDelete)

	adminGroup.Get("/menus", handler.AdminMenusIndex)
	adminGroup.Post("/menus", handler.AdminMenusCreate)
	adminGroup.Post("/menus/:id/delete", handler.AdminMenusDelete)

	adminGroup.Get("/users", handler.AdminUsersIndex)

	// Dynamic page catch-all (MUST be last)
	app.Get("/:slug", handler.Dynamic)

	// Fallback 404 for anything else
	app.Use(handler.NotFound)

	log.Printf("LagariGo listening on http://localhost:%s", cfg.AppPort)
	if err := app.Listen(":" + cfg.AppPort); err != nil {
		log.Fatal(err)
	}
}
