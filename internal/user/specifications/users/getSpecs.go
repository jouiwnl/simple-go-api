package userSpecifications

import (
	"github.com/jouiwnl/simple-go-api/internal/application/commons"
	"github.com/jouiwnl/simple-go-api/internal/user/dto"
	"github.com/jouiwnl/simple-go-api/internal/user/repository"
)

func GetSpecs(candidate *dto.UserDto, repository *repository.UserRepository) *[]commons.Specification[dto.UserDto] {
	return &[]commons.Specification[dto.UserDto]{
		NewVerifyEmptyCpf(candidate),
		NewVerifyEmptyEmail(candidate),
		NewVerifyEmptyName(candidate),
		NewVerifyDuplicatedCpf(candidate, repository),
		NewVerifyDuplicatedEmail(candidate, repository),
	}
}
