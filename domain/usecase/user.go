package usecasedomain

import (
	"context"
	RepoDomain "go-authorization/domain/repository"
)

type UserUsecase interface {
	GetAll(ctx context.Context) ([]RepoDomain.User, error)
	Get(ctx context.Context, email string) (RepoDomain.User, error)
	GetByRole(ctx context.Context, roleName string) ([]RepoDomain.User, error)
	Create(ctx context.Context, user *RepoDomain.User) error
	Update(ctx context.Context, email string, user *RepoDomain.User) error
	Delete(ctx context.Context, email string) error
}
