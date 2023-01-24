package usecase

import (
	"context"
	RepoDomain "go-authorization/domain/repository"
	OrmFactory "go-authorization/internal/repository/orm/factory"
)

type UserUsecase struct {
	orm *OrmFactory.OrmRepository
}

func NewUserUsecase(orm *OrmFactory.OrmRepository) *UserUsecase {
	return &UserUsecase{
		orm: orm,
	}
}

func (rsa *UserUsecase) GetAll(ctx context.Context) ([]RepoDomain.User, error) {
	return rsa.orm.User.GetAll()
}

func (rsa *UserUsecase) Get(ctx context.Context, email string) (RepoDomain.User, error) {
	return rsa.orm.User.Get(email)
}

func (rsa *UserUsecase) GetByRole(ctx context.Context, roleName string) ([]RepoDomain.User, error) {
	return rsa.orm.User.GetByRole(roleName)
}

func (rsa *UserUsecase) Create(ctx context.Context, user *RepoDomain.User) error {
	return rsa.orm.User.Create(user)
}

func (rsa *UserUsecase) Update(ctx context.Context, email string, user *RepoDomain.User) error {
	return rsa.orm.User.Update(email, user)
}

func (rsa *UserUsecase) Delete(ctx context.Context, email string) error {
	return rsa.orm.User.Delete(email)
}
