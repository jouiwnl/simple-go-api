package userSpecifications

import (
	"fmt"
	"github.com/jouiwnl/simple-go-api/internal/user/dto"
	"github.com/jouiwnl/simple-go-api/internal/user/repository"
)

type VerifyDuplicatedCpf struct {
	candidate      *dto.UserDto
	userRepository *repository.UserRepository
}

func NewVerifyDuplicatedCpf(candidate *dto.UserDto, userRepository *repository.UserRepository) *VerifyDuplicatedCpf {
	return &VerifyDuplicatedCpf{
		candidate:      candidate,
		userRepository: userRepository,
	}
}

func (spec VerifyDuplicatedCpf) IsSatisfiedBy(userDto dto.UserDto) (bool, string) {

	exists, _ := spec.userRepository.ExistsUserByCpf(spec.candidate.ID, spec.candidate.Cpf)

	if exists {
		return false, fmt.Sprintf("Já existe um usuário cadastrado com o cpf %s.", *spec.candidate.Cpf)
	}

	return true, ""
}
