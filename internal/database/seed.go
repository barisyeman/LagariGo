package database

import (
	"log"

	"github.com/barisyeman/LagariGo/internal/config"
	"golang.org/x/crypto/bcrypt"
)

// Seed creates the default admin user and demo content on first boot.
func Seed(cfg *config.Config) error {
	var count int64
	DB.Model(&User{}).Count(&count)
	if count == 0 {
		hash, err := bcrypt.GenerateFromPassword([]byte(cfg.AdminPassword), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		admin := User{
			Email:    cfg.AdminEmail,
			Password: string(hash),
			Name:     "Administrator",
			Role:     "admin",
		}
		if err := DB.Create(&admin).Error; err != nil {
			return err
		}
		log.Printf("seeded admin user: %s / %s", cfg.AdminEmail, cfg.AdminPassword)
	}

	DB.Model(&Page{}).Count(&count)
	if count == 0 {
		pages := []Page{
			{Slug: "welcome", Title: "Welcome", Content: "This is a sample page created by LagariGo. Edit it from the admin panel or replace it with your own content.", Published: true},
		}
		DB.Create(&pages)
	}

	DB.Model(&Menu{}).Count(&count)
	if count == 0 {
		menus := []Menu{
			{Label: "Home", URL: "/", Location: "header", Position: 1},
			{Label: "About", URL: "/about-us", Location: "header", Position: 2},
			{Label: "Contact", URL: "/contact", Location: "header", Position: 3},
			{Label: "About", URL: "/about-us", Location: "footer", Position: 1},
			{Label: "Contact", URL: "/contact", Location: "footer", Position: 2},
			{Label: "Sign in", URL: "/login", Location: "footer", Position: 3},
		}
		DB.Create(&menus)
	}
	return nil
}
