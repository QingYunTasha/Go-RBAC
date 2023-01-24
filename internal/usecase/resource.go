package usecase

import (
	OrmFactory "go-authorization/internal/repository/orm/factory"
)

type ResourceApi struct {
	orm *OrmFactory.OrmRepository
}

func NewResourceApi(orm *OrmFactory.OrmRepository) *ResourceApi {
	return &ResourceApi{
		orm: orm,
	}
}
