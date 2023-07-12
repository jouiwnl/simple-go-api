package userSpecifications

import (
	"github.com/jouiwnl/simple-go-api/internal/user/dto"
)

type VerifyEmptyCpf struct {
	candidate *dto.UserDto
}

func NewVerifyEmptyCpf(candidate *dto.UserDto) *VerifyEmptyCpf {
	return &VerifyEmptyCpf{
		candidate: candidate,
	}
}

func (spec VerifyEmptyCpf) IsSatisfiedBy(userDto dto.UserDto) (bool, string) {
	emptyString := ""

	if spec.candidate.Cpf == nil || spec.candidate.Cpf == &emptyString {
		return false, "O cpf do usuário e obrigatório."
	}

	return true, ""
}
