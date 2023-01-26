package repositorydomain

import "gorm.io/gorm"

type operation string

const (
	WRITE operation = "write"
	READ  operation = "read"
)

type Permission struct {
	gorm.Model
	Operation    operation `gorm:"not null"`
	ResourceName string
}

type PermissionRepository interface {
	GetAll() ([]Permission, error)
	GetByResource(resourceName string) ([]Permission, error)
	Create(permission *Permission) error
	Delete(resourceName string, operation string) error
}
