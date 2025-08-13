package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectToDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
