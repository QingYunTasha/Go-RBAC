package repositorydomain

type Permission struct {
	Operation    string `gorm:"primaryKey;not null;default:null"`
	ResourceName string `gorm:"primaryKey;not null;default:null"`
}

type PermissionRepository interface {
	GetAll() ([]Permission, error)
	GetByResource(resourceName string) ([]Permission, error)
	Create(permission *Permission) error
	Delete(resourceName string, operation string) error
}
