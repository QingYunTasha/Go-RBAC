package repository

import (
	RepoDomain "go-authorization/domain/repository"

	"gorm.io/gorm"
)

type RoleOrm struct {
	Db *gorm.DB
}

func NewRoleOrm(db *gorm.DB) *RoleOrm {
	return &RoleOrm{
		Db: db,
	}
}

func (ro *RoleOrm) GetAll() ([]RepoDomain.Role, error) {
	roles := []RepoDomain.Role{}
	err := ro.Db.Find(&roles).Error
	return roles, err
}

func (ro *RoleOrm) Get(name string) (RepoDomain.Role, error) {
	role := RepoDomain.Role{}
	err := ro.Db.Where("Name = ?", name).Take(&role).Error
	return role, err
}

func (ro *RoleOrm) Create(role *RepoDomain.Role) error {
	return ro.Db.Create(role).Error
}

func (ro *RoleOrm) Update(name string, role *RepoDomain.Role) error {
	return ro.Db.Model(&RepoDomain.Role{}).Where("Name = ?", name).Updates(role).Error
}

func (ro *RoleOrm) Delete(name string) error {
	return ro.Db.Where("Name = ?", name).Delete(&RepoDomain.Role{}).Error
}

func (ro *UserOrm) AppendUserAssociation(user *RepoDomain.User) error {
	return ro.Db.Model(&RepoDomain.Role{}).Where("Name = ?", user.RoleName).Association("Users").Append(user)
}

func (ro *UserOrm) DeleteUserAssociation(user *RepoDomain.User) error {
	return ro.Db.Model(&RepoDomain.Role{}).Where("Name = ?", user.RoleName).Association("Users").Delete(user)
}

func (ro *UserOrm) ClearUserAssociation(user *RepoDomain.User) error {
	return ro.Db.Model(&RepoDomain.Role{}).Where("Name = ?", user.RoleName).Association("Users").Clear()
}

func (ro *UserOrm) CountUserAssociation(user *RepoDomain.User) int64 {
	return ro.Db.Model(&RepoDomain.Role{}).Where("Name = ?", user.RoleName).Association("Users").Count()
}
