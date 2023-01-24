package repositorydomain

import "gorm.io/gorm"

type Level struct {
	gorm.Model
	Name string `gorm:"unique, not null"`
}

type LevelRepository interface{}

type LevelUsecase interface{}
