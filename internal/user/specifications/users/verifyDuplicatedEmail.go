package userSpecifications

import (
	"fmt"
	"github.com/jouiwnl/simple-go-api/internal/user/dto"
	"github.com/jouiwnl/simple-go-api/internal/user/repository"
)

type VerifyDuplicatedEmail struct {
	candidate      *dto.UserDto
	userRepository *repository.UserRepository
}

func NewVerifyDuplicatedEmail(candidate *dto.UserDto, userRepository *repository.UserRepository) *VerifyDuplicatedEmail {
	return &VerifyDuplicatedEmail{
		candidate:      candidate,
		userRepository: userRepository,
	}
}

func (spec VerifyDuplicatedEmail) IsSatisfiedBy(userDto dto.UserDto) (bool, string) {

	exists, _ := spec.userRepository.ExistsUserByEmail(spec.candidate.ID, spec.candidate.Email)

	if exists {
		return false, fmt.Sprintf("Já existe um usuário cadastrado com o e-mail %s.", *spec.candidate.Email)
	}

	return true, ""
}
