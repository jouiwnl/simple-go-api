package models

import "time"

type User struct {
	ID        int        `json:"id" gorm:"primarykey"`
	Name      string     `json:"name"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at"`
}

func (u User) GetId() int64 {
	return int64(u.ID)
}
