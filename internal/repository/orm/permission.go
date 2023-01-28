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
	Permissions := []RepoDomain.Permission{}
	err := rso.Db.Find(&Permissions).Error
	return Permissions, err
}

func (rso *PermissionOrm) GetByResource(resourceName string) ([]RepoDomain.Permission, error) {
	permissions := []RepoDomain.Permission{}
	err := rso.Db.Model(&RepoDomain.Resource{}).Where("Name = ?", resourceName).Association("Permissions").Find(permissions)
	return permissions, err
}

func (rso *PermissionOrm) Create(permission *RepoDomain.Permission) error {
	return rso.Db.Model(&RepoDomain.Resource{}).Where("Name = ?", permission.ResourceName).Association("Permissions").Append(map[string]interface{}{
		"Operation":    permission.Operation,
		"ResourceName": permission.ResourceName,
	})
}

func (rso *PermissionOrm) Delete(resourceName string, operation RepoDomain.Operation) error {
	permission := &RepoDomain.Permission{}
	if err := rso.Db.Where("Operation = ? AND ResourceName = ?", operation, resourceName).First(permission).Error; err != nil {
		return err
	}
	return rso.Db.Model(&RepoDomain.Resource{}).Where("Name = ?", resourceName).Association("Permissions").Delete(permission)
}
