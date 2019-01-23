package service

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/tongzcomradez/go-gorm/model"
)

// UserRepo : service for location
type UserRepo interface {
	FindAll(user *model.User) (*model.User, error)
}

// UserRepoGorm is a DAO for location model
type UserRepoGorm struct {
	db *gorm.DB
}

// NewUserRepo instantiates new UserRepo instance
func NewUserRepo(db *gorm.DB) *UserRepoGorm {
	return &UserRepoGorm{db: db}
}

// FindAll create location record
func (repo *UserRepoGorm) FindAll(users *[]model.User) (*[]model.User, error) {
	db := repo.db.Find(users)
	fmt.Println("inininin", users)
	if db.Error != nil {
		return nil, db.Error
	}

	return users, nil
}

// Create create location record
func (repo *UserRepoGorm) Create(user *model.User) (*model.User, error) {
	db := repo.db.Create(user)
	if db.Error != nil {
		return nil, db.Error
	}

	return user, nil
}

// GetProfile abc
func (repo *UserRepoGorm) GetProfile(user *model.User) (*model.User, error) {
	db := repo.db.Preload("Profiles").First(user, 19)

	if db.Error != nil {

		return nil, db.Error
	}

	return user, nil
}
