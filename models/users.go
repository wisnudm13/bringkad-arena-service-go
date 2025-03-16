package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Uuid        string `gorm:"type:char(36)"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.Uuid = uuid.NewString()
	return
}
