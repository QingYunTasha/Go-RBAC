package repositorydomain

import "gorm.io/gorm"

type Resource struct {
	gorm.Model
	Name        string       `gorm:"unique, not null"`
	Permissions []Permission `gorm:"foreignKey:ResourceName;references:Name"`
}

type ResourceRepository interface {
}

type ResourceUsecase interface {
}
