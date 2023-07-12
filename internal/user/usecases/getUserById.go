package usecases

import (
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

func (u *GetUserByIdUseCase) GetUserById(c *gin.Context) {
	id := c.Param("id")

	response, err := u.userRepository.GetUserById(id)

	if err != nil {
		utils.WriteInternalErrorResponse(c, err.Error())
	}

	c.JSON(200, dto.NewUserDtoByEntity(response))
}

/*func GetUserById(c *gin.Context) {
	id := c.Param("id")

	dbContext := initializers.Database.Where("id = ?", id)

	response, err := repository.FindOne[models.User](dbContext)

	if err != nil {
		c.JSON(404, gin.H{
			"statusCode": 404,
			"message":    err.Error(),
		})
		return
	}

	c.JSON(200, response)
}*/
