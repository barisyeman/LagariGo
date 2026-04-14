package handler

import (
	"github.com/a-h/templ"
	"github.com/barisyeman/LagariGo/internal/auth"
	"github.com/barisyeman/LagariGo/internal/database"
	"github.com/barisyeman/LagariGo/views/layouts"
	"github.com/gofiber/fiber/v2"
)

// Render writes a templ component to the Fiber response.
func Render(c *fiber.Ctx, comp templ.Component) error {
	c.Set("Content-Type", "text/html; charset=utf-8")
	return comp.Render(c.Context(), c.Response().BodyWriter())
}

// BuildPageData populates common layout data: menus, current user, CSRF, flash.
func BuildPageData(c *fiber.Ctx, title string) layouts.PageData {
	var headerMenus, footerMenus []database.Menu
	database.DB.Where("location = ?", "header").Order("position asc").Find(&headerMenus)
	database.DB.Where("location = ?", "footer").Order("position asc").Find(&footerMenus)

	user, _ := c.Locals("user").(*database.User)
	csrf, _ := c.Locals("csrf").(string)

	return layouts.PageData{
		Title:        title,
		User:         user,
		HeaderMenus:  headerMenus,
		FooterMenus:  footerMenus,
		CSRFToken:    csrf,
		FlashSuccess: auth.ConsumeFlash(c, "success"),
		FlashError:   auth.ConsumeFlash(c, "error"),
	}
}
