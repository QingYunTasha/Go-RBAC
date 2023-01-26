package usecase

import (
	"context"
	RepoDomain "go-authorization/domain/repository"
	OrmFactory "go-authorization/internal/repository/orm/factory"
)

type LevelUsecase struct {
	orm *OrmFactory.OrmRepository
}

func NewLevelUsecase(orm *OrmFactory.OrmRepository) *LevelUsecase {
	return &LevelUsecase{
		orm: orm,
	}
}

func (rsa *LevelUsecase) GetAll(ctx context.Context) ([]RepoDomain.Level, error) {
	return rsa.orm.Level.GetAll()
}

func (rsa *LevelUsecase) Get(ctx context.Context, name string) (RepoDomain.Level, error) {
	return rsa.orm.Level.Get(name)
}

func (rsa *LevelUsecase) Create(ctx context.Context, Level *RepoDomain.Level) error {
	return rsa.orm.Level.Create(Level)
}

func (rsa *LevelUsecase) Update(ctx context.Context, name string, Level *RepoDomain.Level) error {
	return rsa.orm.Level.Update(name, Level)
}

func (rsa *LevelUsecase) Delete(ctx context.Context, name string) error {
	return rsa.orm.Level.Delete(name)
}
