package repositorydomain

type User struct {
	Name     string
	Email    string `gorm:"primaryKey;not null;default:null"`
	RoleName string
}

type UserRepository interface {
	GetAll() ([]User, error)
	Get(email string) (User, error)
	GetByRole(roleName string) ([]User, error)
	Create(user *User) error
	Update(email string, user *User) error
	Delete(email string) error
}
