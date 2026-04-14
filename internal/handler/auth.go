package handler

import (
	"strings"

	"github.com/barisyeman/LagariGo/internal/auth"
	"github.com/barisyeman/LagariGo/internal/database"
	"github.com/barisyeman/LagariGo/views/pages"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func ShowLogin(c *fiber.Ctx) error {
	return Render(c, pages.Login(BuildPageData(c, "Sign in")))
}

func DoLogin(c *fiber.Ctx) error {
	email := strings.TrimSpace(strings.ToLower(c.FormValue("email")))
	password := c.FormValue("password")

	var u database.User
	if err := database.DB.Where("email = ?", email).First(&u).Error; err != nil {
		auth.Flash(c, "error", "Invalid email or password")
		return c.Redirect("/login")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		auth.Flash(c, "error", "Invalid email or password")
		return c.Redirect("/login")
	}
	if err := auth.Login(c, u.ID); err != nil {
		return err
	}
	if u.IsAdmin() {
		return c.Redirect("/admin")
	}
	return c.Redirect("/")
}

func ShowRegister(c *fiber.Ctx) error {
	return Render(c, pages.Register(BuildPageData(c, "Create account")))
}

func DoRegister(c *fiber.Ctx) error {
	name := strings.TrimSpace(c.FormValue("name"))
	email := strings.TrimSpace(strings.ToLower(c.FormValue("email")))
	password := c.FormValue("password")

	if name == "" || email == "" || len(password) < 6 {
		auth.Flash(c, "error", "Please fill in all fields (password must be at least 6 characters)")
		return c.Redirect("/register")
	}

	var existing database.User
	if err := database.DB.Where("email = ?", email).First(&existing).Error; err == nil {
		auth.Flash(c, "error", "This email is already registered")
		return c.Redirect("/register")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u := database.User{Email: email, Password: string(hash), Name: name, Role: "user"}
	if err := database.DB.Create(&u).Error; err != nil {
		return err
	}
	if err := auth.Login(c, u.ID); err != nil {
		return err
	}
	auth.Flash(c, "success", "Welcome aboard! Your account has been created.")
	return c.Redirect("/")
}

func DoLogout(c *fiber.Ctx) error {
	auth.Logout(c)
	return c.Redirect("/")
}
