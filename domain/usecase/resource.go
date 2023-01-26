package usecasedomain

import (
	"context"
	RepoDomain "go-authorization/domain/repository"
)

type ResourceUsecase interface {
	GetAll(ctx context.Context) ([]RepoDomain.Resource, error)
	Get(ctx context.Context, name string) (RepoDomain.Resource, error)
	Create(ctx context.Context, resource *RepoDomain.Resource) error
	Update(ctx context.Context, name string, resource *RepoDomain.Resource) error
	Delete(ctx context.Context, name string) error
}
