package usecases

import (
	"github.com/gin-gonic/gin"
	specificationHelper "github.com/jouiwnl/simple-go-api/internal/application/helpers"
	"github.com/jouiwnl/simple-go-api/internal/application/utils"
	"github.com/jouiwnl/simple-go-api/internal/user/dto"
	"github.com/jouiwnl/simple-go-api/internal/user/repository"
	userSpecifications "github.com/jouiwnl/simple-go-api/internal/user/specifications/users"
)

type CreateUserUseCase struct {
	userRepository *repository.UserRepository
}

func NewCreateUserUseCase(userRepository *repository.UserRepository) *CreateUserUseCase {
	return &CreateUserUseCase{
		userRepository: userRepository,
	}
}

func (u *CreateUserUseCase) Execute(c *gin.Context) {
	user := &dto.UserDto{}
	err := c.ShouldBind(&user)

	if err != nil {
		utils.WriteInternalErrorResponse(c, err.Error())
		return
	}

	specErrors := specificationHelper.RunSpecs[dto.UserDto](user, userSpecifications.GetSpecs(user, u.userRepository))

	if len(specErrors) > 0 {
		c.JSON(422, gin.H{
			"message": "Erro ao processar usuário",
			"errors":  specErrors,
		})
		return
	}

	created, errorDb := u.userRepository.CreateUser(user)

	if errorDb != nil {
		utils.WriteInternalErrorResponse(c, errorDb.Error())
		return
	}

	c.JSON(200, gin.H{"id": created})
}
