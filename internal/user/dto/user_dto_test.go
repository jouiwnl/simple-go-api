package dto

import (
	"github.com/google/uuid"
	"github.com/jouiwnl/simple-go-api/internal/config/database/sqlc"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserDto(t *testing.T) {
	t.Run("should instantiate a new user dto by sqlc entity", func(t *testing.T) {
		sqlcEntity := sqlc.User{
			ID:   uuid.New().String(),
			Name: "nome",
		}

		userDto := NewUserDtoByEntity(&sqlcEntity)

		assert.Equal(t, sqlcEntity.Name, *userDto.Name)
		assert.Equal(t, sqlcEntity.ID, *userDto.ID)
	})
}
