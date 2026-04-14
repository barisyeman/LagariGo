package database

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey"`
	Email     string    `gorm:"uniqueIndex;not null;size:191"`
	Password  string    `gorm:"not null"`
	Name      string    `gorm:"size:120"`
	Role      string    `gorm:"size:20;default:user"` // user | admin
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) IsAdmin() bool { return u.Role == "admin" }

type Page struct {
	ID        uint      `gorm:"primaryKey"`
	Slug      string    `gorm:"uniqueIndex;not null;size:191"`
	Title     string    `gorm:"not null;size:191"`
	Content   string    `gorm:"type:text"`
	Published bool      `gorm:"default:true"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Menu struct {
	ID       uint   `gorm:"primaryKey"`
	Label    string `gorm:"not null;size:120"`
	URL      string `gorm:"not null;size:255"`
	Location string `gorm:"size:20;default:footer"` // header | footer
	Position int    `gorm:"default:0"`
}

type Setting struct {
	ID    uint   `gorm:"primaryKey"`
	Key   string `gorm:"uniqueIndex;not null;size:120"`
	Value string `gorm:"type:text"`
}
