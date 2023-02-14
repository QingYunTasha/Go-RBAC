package factory

import (
	RepoDomain "go-authorization/domain/repository"
	Repository "go-authorization/internal/repository/database/sql"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitDb(dsn string) (*gorm.DB, error) {
	return gorm.Open(postgres.Open(dsn), &gorm.Config{
		// logger all sql
		//Logger: logger.Default.LogMode(logger.Info),
	})
}

type OrmRepository struct {
	Level      RepoDomain.LevelRepository
	Permission RepoDomain.PermissionRepository
	Resource   RepoDomain.ResourceRepository
	Role       RepoDomain.RoleRepository
	User       RepoDomain.UserRepository
}

func InitOrmRepository(db *gorm.DB) (*OrmRepository, error) {
	if err := db.AutoMigrate(&RepoDomain.Resource{}, &RepoDomain.Level{}, &RepoDomain.Permission{}, &RepoDomain.Role{}, &RepoDomain.User{}); err != nil {
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
