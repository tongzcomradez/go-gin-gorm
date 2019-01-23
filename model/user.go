package model

import (
	"github.com/jinzhu/gorm"
)

// User createt table users
type User struct {
	gorm.Model
	Name     string    `json:"name"`
	Profiles []Profile `gorm:"many2many:user_profiles;"`
}
