package auth

import (
	"time"

	"github.com/barisyeman/LagariGo/internal/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

var Store *session.Store

func Init() {
	Store = session.New(session.Config{
		Expiration:     24 * time.Hour,
		KeyLookup:      "cookie:lagarigo_session",
		CookieHTTPOnly: true,
		CookieSameSite: "Lax",
	})
}

func Login(c *fiber.Ctx, userID uint) error {
	sess, err := Store.Get(c)
	if err != nil {
		return err
	}
	sess.Set("user_id", userID)
	return sess.Save()
}

func Logout(c *fiber.Ctx) error {
	sess, err := Store.Get(c)
	if err != nil {
		return err
	}
	return sess.Destroy()
}

func CurrentUser(c *fiber.Ctx) *database.User {
	sess, err := Store.Get(c)
	if err != nil {
		return nil
	}
	id, ok := sess.Get("user_id").(uint)
	if !ok {
		return nil
	}
	var u database.User
	if err := database.DB.First(&u, id).Error; err != nil {
		return nil
	}
	return &u
}

// Flash stores a one-shot message (consumed on next read).
func Flash(c *fiber.Ctx, key, msg string) {
	sess, err := Store.Get(c)
	if err != nil {
		return
	}
	sess.Set("flash_"+key, msg)
	sess.Save()
}

func ConsumeFlash(c *fiber.Ctx, key string) string {
	sess, err := Store.Get(c)
	if err != nil {
		return ""
	}
	v, _ := sess.Get("flash_" + key).(string)
	if v != "" {
		sess.Delete("flash_" + key)
		sess.Save()
	}
	return v
}
