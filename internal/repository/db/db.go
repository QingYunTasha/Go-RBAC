package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitDb(dsn string) (*gorm.DB, error) {
	return gorm.Open(postgres.Open(dsn))
}
