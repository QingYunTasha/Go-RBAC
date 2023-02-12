package repositorydomain

type Role struct {
	Name        string        `gorm:"primaryKey;not null;default:null"`
	Users       []User        `gorm:"foreignKey:RoleName;references:Name;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Permissions []*Permission `gorm:"many2many:role_permissions;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

type RoleRepository interface {
	GetAll() ([]Role, error)
	Get(name string) (Role, error)
	GetByPermission(permission *Permission) ([]Role, error)
	GetByUser(user *User) (Role, error)
	Create(role *Role) error
	Update(name string, role *Role) error
	Delete(name string) error
}
