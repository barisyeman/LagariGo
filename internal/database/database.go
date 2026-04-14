package database

import (
	"fmt"
	"log"

	"github.com/barisyeman/LagariGo/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(cfg *config.Config) error {
	var dialector gorm.Dialector
	switch cfg.DBDriver {
	case "mysql":
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)
		dialector = mysql.Open(dsn)
	default:
		dialector = sqlite.Open(cfg.DBSQLitePath)
	}

	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		return err
	}
	DB = db
	log.Printf("database connected (%s)", cfg.DBDriver)
	return nil
}

func Migrate() error {
	return DB.AutoMigrate(&User{}, &Page{}, &Menu{}, &Setting{})
}
