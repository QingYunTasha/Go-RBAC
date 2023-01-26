package repositorydomain

import "gorm.io/gorm"

type Level struct {
	gorm.Model
	Name string `gorm:"unique, not null"`
}

type LevelRepository interface {
	GetAll() ([]Level, error)
	Get(name string) (Level, error)
	Create(level *Level) error
	Update(name string, level *Level) error
	Delete(name string) error
}
