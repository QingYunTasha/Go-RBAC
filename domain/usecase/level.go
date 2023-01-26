package usecasedomain

import (
	"context"
	RepoDomain "go-authorization/domain/repository"
)

type LevelUsecase interface {
	GetAll(ctx context.Context) ([]RepoDomain.Level, error)
	Get(ctx context.Context, name string) (RepoDomain.Level, error)
	Create(ctx context.Context, level *RepoDomain.Level) error
	Update(ctx context.Context, name string, level *RepoDomain.Level) error
	Delete(ctx context.Context, name string) error
}
