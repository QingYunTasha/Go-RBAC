package usecase

import (
	"context"
	UsecaseDomain "go-authorization/domain/usecase"
	OrmFactory "go-authorization/internal/repository/database/factory"
)

type CoreUsecase struct {
	orm *OrmFactory.OrmRepository
}

func NewCoreUsecase(orm *OrmFactory.OrmRepository) *CoreUsecase {
	return &CoreUsecase{
		orm: orm,
	}
}

func (cra *CoreUsecase) CheckPermission(ctx context.Context, userEmail string, action string, resource string) (UsecaseDomain.HasPermission, error) {
	user, err := cra.orm.User.Get(userEmail)
	if err != nil {
		return false, err
	}

	role, err := cra.orm.Role.GetByUser(&user)
	if err != nil {
		return false, err
	}

	permissions, err := cra.orm.Permission.GetByRole(&role)
	if err != nil {
		return false, err
	}

	for _, permission := range permissions {
		if permission.Action == action && permission.ResourceName == resource {
			return true, nil
		}
	}

	return false, nil
}
