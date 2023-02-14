package usecase

import (
	"context"
	RepoDomain "go-authorization/domain/repository"
	OrmFactory "go-authorization/internal/repository/database/factory"
)

type ResourceUsecase struct {
	orm *OrmFactory.OrmRepository
}

func NewResourceUsecase(orm *OrmFactory.OrmRepository) *ResourceUsecase {
	return &ResourceUsecase{
		orm: orm,
	}
}

func (rsa *ResourceUsecase) GetAll(ctx context.Context) ([]RepoDomain.Resource, error) {
	return rsa.orm.Resource.GetAll()
}

func (rsa *ResourceUsecase) Get(ctx context.Context, name string) (RepoDomain.Resource, error) {
	return rsa.orm.Resource.Get(name)
}

func (rsa *ResourceUsecase) Create(ctx context.Context, Resource *RepoDomain.Resource) error {
	return rsa.orm.Resource.Create(Resource)
}

func (rsa *ResourceUsecase) Update(ctx context.Context, name string, Resource *RepoDomain.Resource) error {
	return rsa.orm.Resource.Update(name, Resource)
}

func (rsa *ResourceUsecase) Delete(ctx context.Context, name string) error {
	return rsa.orm.Resource.Delete(name)
}
