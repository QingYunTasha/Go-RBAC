package factory

import (
	RepoDomain "go-authorization/domain/repository"
	Repository "go-authorization/internal/repository/orm"

	"gorm.io/gorm"
)

type OrmRepository struct {
	Level      RepoDomain.LevelRepository
	Permission RepoDomain.PermissionRepository
	Resource   RepoDomain.ResourceRepository
	Role       RepoDomain.RoleRepository
	User       RepoDomain.UserRepository
}

func InitOrmRepository(db *gorm.DB) (*OrmRepository, error) {
	if err := db.AutoMigrate(&RepoDomain.Resource{}, &RepoDomain.Level{}); err != nil {
		return nil, err
	}

	return &OrmRepository{
		Level:      Repository.NewLevelOrm(db),
		Permission: Repository.NewPermissionOrm(db),
		Resource:   Repository.NewResourceOrm(db),
		Role:       Repository.NewRoleOrm(db),
		User:       Repository.NewUserOrm(db),
	}, nil
}
