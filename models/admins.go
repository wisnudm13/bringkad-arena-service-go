package models

import (
	"bringkad-arena-service-go/utils"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model
	Uuid     string `gorm:"type:char(36);uniqueIndex"`
	Username string `gorm:"uniqueIndex;not null" json:"username"`
	Email    string `gorm:"uniqueIndex;not null" json:"email"`
	IsActive bool   `gorm:"default:true" json:"is_active"`
	Password string `gorm:"uniqueIndex;not null" json:"password"`
}

func (a *Admin) BeforeCreate(tx *gorm.DB) (err error) {
	// auto generate uuid as string
	a.Uuid = uuid.NewString()

	// hash password
	hashedPassword, err := utils.HashPassword(a.Password)
	if err != nil {
		return err
	}

	a.Password = hashedPassword

	return
}
