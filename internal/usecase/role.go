package usecase

import (
	"context"
	RepoDomain "go-authorization/domain/repository"
	OrmFactory "go-authorization/internal/repository/orm/factory"
)

type RoleUsecase struct {
	orm *OrmFactory.OrmRepository
}

func NewRoleUsecase(orm *OrmFactory.OrmRepository) *RoleUsecase {
	return &RoleUsecase{
		orm: orm,
	}
}

func (rsa *RoleUsecase) GetAll(ctx context.Context) ([]RepoDomain.Role, error) {
	return rsa.orm.Role.GetAll()
}

func (rsa *RoleUsecase) Get(ctx context.Context, name string) (RepoDomain.Role, error) {
	return rsa.orm.Role.Get(name)
}

func (rsa *RoleUsecase) Create(ctx context.Context, Role *RepoDomain.Role) error {
	return rsa.orm.Role.Create(Role)
}

func (rsa *RoleUsecase) Update(ctx context.Context, name string, Role *RepoDomain.Role) error {
	return rsa.orm.Role.Update(name, Role)
}

func (rsa *RoleUsecase) Delete(ctx context.Context, name string) error {
	return rsa.orm.Role.Delete(name)
}
