package usecases

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jouiwnl/simple-go-api/internal/application/utils"
	"github.com/jouiwnl/simple-go-api/internal/user/repository"
)

type DeleteUserCase struct {
	userRepository *repository.UserRepository
}

func NewDeleteUserCase(userRepository *repository.UserRepository) *DeleteUserCase {
	return &DeleteUserCase{
		userRepository: userRepository,
	}
}

func (u *DeleteUserCase) Execute(c *gin.Context) {
	id := c.Param("id")

	err := u.userRepository.DeleteUser(&id)

	if err != nil {
		utils.WriteInternalErrorResponse(c, err.Error())
		return
	}

	c.JSON(200, http.NoBody)
}
