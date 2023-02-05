package repositorydomain

type Permission struct {
	Operation    string `gorm:"primaryKey"`
	ResourceName string `gorm:"primaryKey"`
}

type PermissionRepository interface {
	GetAll() ([]Permission, error)
	GetByResource(resourceName string) ([]Permission, error)
	Create(permission *Permission) error
	Delete(resourceName string, operation string) error
}
