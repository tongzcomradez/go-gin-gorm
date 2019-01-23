package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// Profile belongs to `User`, `UserID` is the foreign key
type Profile struct {
	gorm.Model

	Name string
}

// BeforeCreate belongs to `User`, `UserID` is the foreign key
func (p *Profile) BeforeCreate(scope *gorm.Scope) (err error) {
	fmt.Println("Before Create")
	return
}

// AfterCreate belongs to `User`, `UserID` is the foreign key
func (p *Profile) AfterCreate(scope *gorm.Scope) (err error) {
	fmt.Println("After Create")
	return
}
