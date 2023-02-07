package repositorydomain

type Role struct {
	Name        string        `gorm:"primaryKey;not null;default:null"`
	Users       []User        `gorm:"foreignKey:RoleName;references:Name"`
	Permissions []*Permission `gorm:"many2many:user_permissions;"`
}

type RoleRepository interface {
	GetAll() ([]Role, error)
	Get(name string) (Role, error)
	GetByPermission(permission *Permission) ([]Role, error)
	Create(role *Role) error
	Update(name string, role *Role) error
	Delete(name string) error
}
