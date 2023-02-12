package repository

import (
	RepoDomain "go-authorization/domain/repository"

	"gorm.io/gorm"
)

type PermissionOrm struct {
	Db *gorm.DB
}

func NewPermissionOrm(db *gorm.DB) *PermissionOrm {
	return &PermissionOrm{
		Db: db,
	}
}

func (rso *PermissionOrm) GetAll() ([]RepoDomain.Permission, error) {
	permissions := []RepoDomain.Permission{}
	err := rso.Db.Find(&permissions).Error
	return permissions, err
}

func (rso *PermissionOrm) GetByResource(resourceName string) ([]RepoDomain.Permission, error) {
	permissions := []RepoDomain.Permission{}
	err := rso.Db.Model(&RepoDomain.Resource{Name: resourceName}).Association("Permissions").Find(&permissions)
	return permissions, err
}

func (rso *PermissionOrm) GetByRole(role *RepoDomain.Role) ([]RepoDomain.Permission, error) {
	var permissions []RepoDomain.Permission
	err := rso.Db.Model(role).Association("Permissions").Find(&permissions)
	return permissions, err
}

func (rso *PermissionOrm) Create(permission *RepoDomain.Permission) error {
	return rso.Db.Create(&permission).Error
}

func (rso *PermissionOrm) Delete(resourceName string, action string) error {
	permission := RepoDomain.Permission{
		ResourceName: resourceName,
		Action:       action,
	}

	if err := rso.Db.Take(&permission).Error; err != nil {
		return err
	}
	return rso.Db.Delete(&permission).Error
}
