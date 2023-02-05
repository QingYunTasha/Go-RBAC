package repository

import (
	RepoDomain "go-authorization/domain/repository"

	"gorm.io/gorm"
)

type LevelOrm struct {
	Db *gorm.DB
}

func NewLevelOrm(db *gorm.DB) *LevelOrm {
	return &LevelOrm{
		Db: db,
	}
}

func (lvo *LevelOrm) GetAll() ([]RepoDomain.Level, error) {
	Levels := []RepoDomain.Level{}
	err := lvo.Db.Find(&Levels).Error
	return Levels, err
}

func (lvo *LevelOrm) Get(name string) (RepoDomain.Level, error) {
	Level := RepoDomain.Level{}
	err := lvo.Db.Where("Name = ?", name).First(&Level).Error
	return Level, err
}

func (lvo *LevelOrm) Create(level *RepoDomain.Level) error {
	return lvo.Db.Model(&RepoDomain.Level{}).Create(map[string]interface{}{
		"Name": level.Name,
	}).Error
}

func (lvo *LevelOrm) Update(name string, level *RepoDomain.Level) error {
	return lvo.Db.Model(&RepoDomain.Level{}).Where("Name = ?", name).Updates(map[string]interface{}{
		"Name": level.Name,
	}).Error
}

func (lvo *LevelOrm) Delete(name string) error {
	return lvo.Db.Where("Name = ?", name).Delete(&RepoDomain.Level{}).Error
}
