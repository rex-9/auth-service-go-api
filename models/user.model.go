package models

import (
	"gorm.io/gorm"
)

type Role string

func (r Role) String() {
	panic("unimplemented")
}

const (
	AdminRole Role = "admin"
	UserRole  Role = "user"
)

type User struct {
	gorm.Model        //this inculue id,creation,edit,delete dateTme
	Email      string `gorm:"unique"`
	Password   string
	ISVerified bool
	UserName   string
	Role       Role `gorm:"default:user"` // default role is user
}

func (user *User) BeforeCreate(tx *gorm.DB) error {
	user.ISVerified = false
	return nil
}
