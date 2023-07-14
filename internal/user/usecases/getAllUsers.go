package usecases

import (
	"strconv"

	"github.com/jouiwnl/simple-go-api/internal/application/commons"
	"github.com/jouiwnl/simple-go-api/internal/application/utils"
	"github.com/jouiwnl/simple-go-api/internal/user/dto"
	"github.com/jouiwnl/simple-go-api/internal/user/repository"

	"github.com/gin-gonic/gin"
)

type GetAllUsersUseCase struct {
	userRepository *repository.UserRepository
}

func NewGetAllUsersUseCase(userRepository *repository.UserRepository) *GetAllUsersUseCase {
	return &GetAllUsersUseCase{
		userRepository: userRepository,
	}
}

func (u *GetAllUsersUseCase) GetUsers(c *gin.Context) {
	var response []dto.UserDto
	users := u.userRepository.GetUsers()

	for _, user := range users {
		response = append(response, *dto.NewUserDtoByEntity(user))
	}

	c.JSON(200, response)
}

func (u *GetAllUsersUseCase) Execute(c *gin.Context) {
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))

	pageable := commons.Pageable{Offset: offset, Limit: limit}

	paginatedUsers, err := u.userRepository.GetPaginatedUsers(pageable)

	if err != nil {
		utils.WriteInternalErrorResponse(c, err.Error())
		return
	}

	c.JSON(200, paginatedUsers)
}
