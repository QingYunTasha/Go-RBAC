package usecase

import (
	"context"
	RepoDomain "go-authorization/domain/repository"
	OrmFactory "go-authorization/internal/repository/orm/factory"
)

type PermissionUsecase struct {
	orm *OrmFactory.OrmRepository
}

func NewPermissionUsecase(orm *OrmFactory.OrmRepository) *PermissionUsecase {
	return &PermissionUsecase{
		orm: orm,
	}
}

func (rsa *PermissionUsecase) GetAll(ctx context.Context) ([]RepoDomain.Permission, error) {
	return rsa.orm.Permission.GetAll()
}

func (rsa *PermissionUsecase) GetByResource(ctx context.Context, resourceName string) ([]RepoDomain.Permission, error) {
	return rsa.orm.Permission.GetByResource(resourceName)
}

func (rsa *PermissionUsecase) GetByRole(ctx context.Context, role *RepoDomain.Role) ([]RepoDomain.Permission, error) {
	return rsa.orm.Permission.GetByRole(role)
}

func (rsa *PermissionUsecase) Create(ctx context.Context, Permission *RepoDomain.Permission) error {
	return rsa.orm.Permission.Create(Permission)
}

func (rsa *PermissionUsecase) Delete(ctx context.Context, resourceName string, operation string) error {
	return rsa.orm.Permission.Delete(resourceName, operation)
}
