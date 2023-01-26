package repositorydomain

import "gorm.io/gorm"

type Operation string

const (
	WRITE Operation = "write"
	READ  Operation = "read"
)

type Permission struct {
	gorm.Model
	Operation    Operation `gorm:"not null"`
	ResourceName string
}

type PermissionRepository interface {
	GetAll() ([]Permission, error)
	GetByResource(resourceName string) ([]Permission, error)
	Create(permission *Permission) error
	Delete(resourceName string, operation string) error
}
