package db

import (
	"github.com/zxstrike/cafe-pos/menu-services/config"
	"github.com/zxstrike/cafe-pos/menu-services/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db gorm.DB

func MySQLConnect(config *config.MySQLConfig) (*gorm.DB, error) {

	dsn := config.Username + ":" + config.Password + "@tcp(" + config.Host + ":" + config.Port + ")/" + config.Database + "?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if db == nil {
		return nil, err
	}

	migration(db)

	Db = *db

	return db, nil
}

func migration(db *gorm.DB) {
	// Migrate the schema
	db.AutoMigrate(
		&models.Menu{},
		&models.MenuHistory{},
		&models.Category{},
	)
}
