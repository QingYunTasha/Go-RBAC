package repositorydomain

type Resource struct {
	Name        string       `gorm:"primaryKey;not null;default:null"`
	Permissions []Permission `gorm:"foreignKey:ResourceName;references:Name;constraint:OnDelete:CASCADE"`
}

type ResourceRepository interface {
	GetAll() ([]Resource, error)
	Get(name string) (Resource, error)
	Create(resource *Resource) error
	Update(name string, resource *Resource) error
	Delete(name string) error
}
