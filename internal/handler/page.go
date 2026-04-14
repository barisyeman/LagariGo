package handler

import (
	"github.com/barisyeman/LagariGo/internal/database"
	"github.com/barisyeman/LagariGo/views/pages"
	"github.com/gofiber/fiber/v2"
)

// ReservedSlugs cannot be used by dynamic pages (collide with static routes).
var ReservedSlugs = map[string]bool{
	"":          true,
	"about-us":  true,
	"contact":   true,
	"login":     true,
	"register":  true,
	"logout":    true,
	"admin":     true,
	"assets":    true,
}

func Home(c *fiber.Ctx) error {
	return Render(c, pages.Home(BuildPageData(c, "Home")))
}

func About(c *fiber.Ctx) error {
	return Render(c, pages.About(BuildPageData(c, "About")))
}

func Contact(c *fiber.Ctx) error {
	return Render(c, pages.Contact(BuildPageData(c, "Contact")))
}

// Dynamic resolves /:slug to a Page record. Returns 404 templ page if missing.
func Dynamic(c *fiber.Ctx) error {
	slug := c.Params("slug")
	if ReservedSlugs[slug] {
		return notFound(c)
	}
	var p database.Page
	if err := database.DB.Where("slug = ? AND published = ?", slug, true).First(&p).Error; err != nil {
		return notFound(c)
	}
	return Render(c, pages.Dynamic(BuildPageData(c, p.Title), p))
}

func notFound(c *fiber.Ctx) error {
	c.Status(fiber.StatusNotFound)
	return Render(c, pages.NotFound(BuildPageData(c, "Not found")))
}

// NotFound is the global 404 handler.
func NotFound(c *fiber.Ctx) error { return notFound(c) }
