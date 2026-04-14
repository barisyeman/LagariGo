package handler

import (
	"strconv"
	"strings"

	"github.com/barisyeman/LagariGo/internal/auth"
	"github.com/barisyeman/LagariGo/internal/database"
	"github.com/barisyeman/LagariGo/views/pages/admin"
	"github.com/gofiber/fiber/v2"
	"github.com/gosimple/slug"
)

func AdminDashboard(c *fiber.Ctx) error {
	stats := map[string]int64{}
	var n int64
	database.DB.Model(&database.Page{}).Count(&n); stats["pages"] = n
	database.DB.Model(&database.Menu{}).Count(&n); stats["menus"] = n
	database.DB.Model(&database.User{}).Count(&n); stats["users"] = n
	return Render(c, admin.Dashboard(BuildPageData(c, "Dashboard"), stats))
}

// --- Pages CRUD ---

func AdminPagesIndex(c *fiber.Ctx) error {
	var pages []database.Page
	database.DB.Order("created_at desc").Find(&pages)
	return Render(c, admin.PagesList(BuildPageData(c, "Pages"), pages))
}

func AdminPagesNew(c *fiber.Ctx) error {
	return Render(c, admin.PageForm(BuildPageData(c, "New page"), database.Page{Published: true}, true))
}

func AdminPagesEdit(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var p database.Page
	if err := database.DB.First(&p, id).Error; err != nil {
		return c.Redirect("/admin/pages")
	}
	return Render(c, admin.PageForm(BuildPageData(c, "Edit page"), p, false))
}

func AdminPagesCreate(c *fiber.Ctx) error {
	p := pageFromForm(c, database.Page{})
	if ReservedSlugs[p.Slug] {
		auth.Flash(c, "error", "This slug is reserved, please choose another")
		return c.Redirect("/admin/pages/new")
	}
	if err := database.DB.Create(&p).Error; err != nil {
		auth.Flash(c, "error", "Could not create page: "+err.Error())
		return c.Redirect("/admin/pages/new")
	}
	auth.Flash(c, "success", "Page created")
	return c.Redirect("/admin/pages")
}

func AdminPagesUpdate(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var p database.Page
	if err := database.DB.First(&p, id).Error; err != nil {
		return c.Redirect("/admin/pages")
	}
	updated := pageFromForm(c, p)
	if ReservedSlugs[updated.Slug] {
		auth.Flash(c, "error", "This slug is reserved")
		return c.Redirect("/admin/pages/" + c.Params("id") + "/edit")
	}
	database.DB.Save(&updated)
	auth.Flash(c, "success", "Page updated")
	return c.Redirect("/admin/pages")
}

func AdminPagesDelete(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	database.DB.Delete(&database.Page{}, id)
	auth.Flash(c, "success", "Page deleted")
	return c.Redirect("/admin/pages")
}

func pageFromForm(c *fiber.Ctx, base database.Page) database.Page {
	base.Title = strings.TrimSpace(c.FormValue("title"))
	s := strings.TrimSpace(c.FormValue("slug"))
	if s == "" {
		s = slug.Make(base.Title)
	}
	base.Slug = strings.ToLower(s)
	base.Content = c.FormValue("content")
	base.Published = c.FormValue("published") == "1"
	return base
}

// --- Menus CRUD ---

func AdminMenusIndex(c *fiber.Ctx) error {
	var menus []database.Menu
	database.DB.Order("location asc, position asc").Find(&menus)
	return Render(c, admin.MenusList(BuildPageData(c, "Menus"), menus))
}

func AdminMenusCreate(c *fiber.Ctx) error {
	pos, _ := strconv.Atoi(c.FormValue("position"))
	loc := c.FormValue("location")
	if loc != "header" && loc != "footer" {
		loc = "footer"
	}
	m := database.Menu{
		Label:    strings.TrimSpace(c.FormValue("label")),
		URL:      strings.TrimSpace(c.FormValue("url")),
		Location: loc,
		Position: pos,
	}
	if m.Label == "" || m.URL == "" {
		auth.Flash(c, "error", "Label and URL are required")
		return c.Redirect("/admin/menus")
	}
	database.DB.Create(&m)
	auth.Flash(c, "success", "Menu link added")
	return c.Redirect("/admin/menus")
}

func AdminMenusDelete(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	database.DB.Delete(&database.Menu{}, id)
	auth.Flash(c, "success", "Link deleted")
	return c.Redirect("/admin/menus")
}

// --- Users (read-only) ---

func AdminUsersIndex(c *fiber.Ctx) error {
	var users []database.User
	database.DB.Order("created_at desc").Find(&users)
	return Render(c, admin.UsersList(BuildPageData(c, "Users"), users))
}
