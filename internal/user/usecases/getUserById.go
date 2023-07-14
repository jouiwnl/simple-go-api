package usecases

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/jouiwnl/simple-go-api/internal/application/utils"
	"github.com/jouiwnl/simple-go-api/internal/user/dto"
	"github.com/jouiwnl/simple-go-api/internal/user/repository"
)

type GetUserByIdUseCase struct {
	userRepository *repository.UserRepository
}

func NewGetUserByIdUseCase(userRepository *repository.UserRepository) *GetUserByIdUseCase {
	return &GetUserByIdUseCase{
		userRepository: userRepository,
	}
}

func (u *GetUserByIdUseCase) Execute(c *gin.Context) {
	id := c.Param("id")

	response, err := u.userRepository.GetUserById(id)

	switch err {
	case sql.ErrNoRows:
		utils.WriteNotFoundErrorResponse(c)
	case nil:
		c.JSON(200, dto.NewUserDtoByEntity(response))
	default:
		utils.WriteInternalErrorResponse(c, err.Error())
	}
}
