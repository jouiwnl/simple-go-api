package usecases

import (
	"github.com/gin-gonic/gin"
	specificationHelper "github.com/jouiwnl/simple-go-api/internal/application/helpers"
	"github.com/jouiwnl/simple-go-api/internal/application/utils"
	"github.com/jouiwnl/simple-go-api/internal/user/dto"
	"github.com/jouiwnl/simple-go-api/internal/user/repository"
	userSpecifications "github.com/jouiwnl/simple-go-api/internal/user/specifications/users"
)

type UpdateUserUseCase struct {
	userRepository *repository.UserRepository
}

func NewUpdateUserUseCase(userRepository *repository.UserRepository) *UpdateUserUseCase {
	return &UpdateUserUseCase{
		userRepository: userRepository,
	}
}

func (u *UpdateUserUseCase) UpdateUser(c *gin.Context) {
	id := c.Param("id")

	user := &dto.UserDto{}
	err := c.BindJSON(&user)

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

	updated, errorDb := u.userRepository.UpdateUser(user, id)

	if errorDb != nil {
		utils.WriteInternalErrorResponse(c, errorDb.Error())
		return
	}

	c.JSON(200, gin.H{"id": updated})
}

/*func UpdateUser(c *gin.Context) {
	id := c.Param("id")

	user := models.User{}
	c.BindJSON(&user)

	database := initializers.Database

	errors := specificationHelper.RunSpecs[models.User](user, userSpecifications.GetSpecs(database))

	if len(errors) > 0 {
		c.JSON(422, gin.H{
			"message": "Erro ao processar entidade",
			"errors":  errors,
		})
		return
	}

	c.JSON(200, repository.Save[models.User](user, id, database))
}*/
