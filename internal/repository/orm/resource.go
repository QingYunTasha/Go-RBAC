package repository

import (
	RepoDomain "go-authorization/domain/repository"

	"gorm.io/gorm"
)

type ResourceOrm struct {
	Db *gorm.DB
}

func NewResourceOrm(db *gorm.DB) *ResourceOrm {
	return &ResourceOrm{
		Db: db,
	}
}

func (rso *ResourceOrm) GetAll() ([]RepoDomain.Resource, error) {
	resources := []RepoDomain.Resource{}
	err := rso.Db.Find(&resources).Error
	return resources, err
}

func (rso *ResourceOrm) Get(name string) (RepoDomain.Resource, error) {
	resource := RepoDomain.Resource{}
	err := rso.Db.Where("Name = ?", name).First(&resource).Error
	return resource, err
}

func (rso *ResourceOrm) Create(resource *RepoDomain.Resource) error {
	return rso.Db.Model(&RepoDomain.Resource{}).Create(map[string]interface{}{
		"Name": resource.Name,
	}).Error
}

func (rso *ResourceOrm) Update(name string, resource *RepoDomain.Resource) error {
	return rso.Db.Model(&RepoDomain.Resource{}).Where("Name = ?", name).Updates(map[string]interface{}{
		"Name": resource.Name,
	}).Error
}

func (rso *ResourceOrm) Delete(name string) error {
	return rso.Db.Where("Name = ?", name).Delete(&RepoDomain.Resource{}).Error
}
