package repositorydomain

import "gorm.io/gorm"

type Resource struct {
	gorm.Model
	Name        string       `gorm:"unique;not null"`
	Permissions []Permission `gorm:"foreignKey:ResourceName;references:Name"`
}

type ResourceRepository interface {
	GetAll() ([]Resource, error)
	Get(name string) (Resource, error)
	Create(resource *Resource) error
	Update(name string, resource *Resource) error
	Delete(name string) error
}
