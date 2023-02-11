package repository

import (
	RepoDomain "go-authorization/domain/repository"

	"gorm.io/gorm"
)

type UserOrm struct {
	Db *gorm.DB
}

func NewUserOrm(db *gorm.DB) *UserOrm {
	return &UserOrm{
		Db: db,
	}
}

func (uso *UserOrm) GetAll() ([]RepoDomain.User, error) {
	users := []RepoDomain.User{}
	err := uso.Db.Find(&users).Error
	return users, err
}

func (uso *UserOrm) Get(email string) (RepoDomain.User, error) {
	user := RepoDomain.User{
		Email: email,
	}
	err := uso.Db.Take(&user).Error
	return user, err
}

func (uso *UserOrm) GetByRole(roleName string) ([]RepoDomain.User, error) {
	users := []RepoDomain.User{}
	err := uso.Db.Model(&RepoDomain.Role{Name: roleName}).Association("Users").Find(users)
	return users, err
}

func (uso *UserOrm) Create(user *RepoDomain.User) error {
	return uso.Db.Create(user).Error
}

func (uso *UserOrm) Update(email string, user *RepoDomain.User) error {
	oldUser, err := uso.Get(email)
	if err != nil {
		return err
	}
	return uso.Db.Model(&oldUser).Updates(&user).Error
}

func (uso *UserOrm) Delete(email string) error {
	user, err := uso.Get(email)
	if err != nil {
		return err
	}

	return uso.Db.Delete(&user).Error
}
