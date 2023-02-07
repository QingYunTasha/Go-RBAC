package usecasedomain

import (
	"context"
	RepoDomain "go-authorization/domain/repository"
)

type RoleUsecase interface {
	GetAll(ctx context.Context) ([]RepoDomain.Role, error)
	Get(ctx context.Context, name string) (RepoDomain.Role, error)
	GetByPermission(ctx context.Context, permission *RepoDomain.Permission) ([]RepoDomain.Role, error)
	Create(ctx context.Context, role *RepoDomain.Role) error
	Update(ctx context.Context, name string, role *RepoDomain.Role) error
	Delete(ctx context.Context, name string) error
}
