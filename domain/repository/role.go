package repositorydomain

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Name  string `gorm:"unique, not null"`
	Users []Role `gorm:"foreignKey:RoleName;references:Name"`
}

type RoleRepository interface {
	GetAll() ([]Role, error)
	Get(name string) (Role, error)
	Create(role *Role) error
	Update(name string, role *Role) error
	Delete(name string) error
}
