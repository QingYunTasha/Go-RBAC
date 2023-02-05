package repositorydomain

type Role struct {
	Name  string `gorm:"primaryKey"`
	Users []User `gorm:"foreignKey:RoleName;references:Name"`
}

type RoleRepository interface {
	GetAll() ([]Role, error)
	Get(name string) (Role, error)
	Create(role *Role) error
	Update(name string, role *Role) error
	Delete(name string) error
}
