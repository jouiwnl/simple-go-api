package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        string     `json:"id" gorm:"primaryKey"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Cpf       string     `json:"cpf"`
	CreatedAt *time.Time `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt *time.Time `json:"updatedAt" gorm:"column:updated_at"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.ID = uuid.NewString()
	return
}
