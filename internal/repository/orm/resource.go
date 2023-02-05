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
	err := rso.Db.Where("name = ?", name).Take(&resource).Error
	return resource, err
}

func (rso *ResourceOrm) Create(resource *RepoDomain.Resource) error {
	return rso.Db.Create(&resource).Error
}

func (rso *ResourceOrm) Update(name string, resource *RepoDomain.Resource) error {
	oldResource, err := rso.Get(name)
	if err != nil {
		return err
	}
	return rso.Db.Model(oldResource).Updates(map[string]interface{}{
		"name": resource.Name,
	}).Error
}

func (rso *ResourceOrm) Delete(name string) error {
	resource, err := rso.Get(name)
	if err != nil {
		return err
	}

	return rso.Db.Where("name = ?", resource.Name).Delete(&resource).Error
}
