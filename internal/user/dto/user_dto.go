package dto

import (
	"github.com/jouiwnl/simple-go-api/internal/config/database/sqlc"
	"time"
)

type UserDto struct {
	ID        *string    `json:"id"`
	Name      *string    `json:"name"`
	Email     *string    `json:"email"`
	Cpf       *string    `json:"cpf"`
	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}

func NewUserDtoByEntity(user *sqlc.User) *UserDto {
	var d UserDto

	if user.Cpf.Valid {
		d.Cpf = &user.Cpf.String
	}

	if user.Email.Valid {
		d.Email = &user.Email.String
	}

	if user.CreatedAt.Valid {
		d.CreatedAt = &user.CreatedAt.Time
	}

	if user.UpdatedAt.Valid {
		d.UpdatedAt = &user.UpdatedAt.Time
	}

	d.ID = &user.ID
	d.Name = &user.Name

	return &d
}
