package usecasedomain

import "context"

type HasPermission bool

const (
	Yes HasPermission = true
	No  HasPermission = false
)

type CoreUsecase interface {
	CheckPermission(ctx context.Context, userEmail string, action string, resource string) (HasPermission, error)
}
