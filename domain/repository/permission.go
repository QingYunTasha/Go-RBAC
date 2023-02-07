package repositorydomain

type Permission struct {
	Operation    string  `gorm:"primaryKey;not null;default:null"`
	ResourceName string  `gorm:"primaryKey;not null;default:null"`
	Roles        []*Role `gorm:"many2many:user_permissions;"`
}

type PermissionRepository interface {
	GetAll() ([]Permission, error)
	GetByResource(resourceName string) ([]Permission, error)
	GetByRole(role *Role) ([]Permission, error)
	Create(permission *Permission) error
	Delete(resourceName string, operation string) error
}
