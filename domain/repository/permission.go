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

type PermissionRepository interface{}

type PermissionUsecase interface{}
