package userSpecifications

import "github.com/jouiwnl/simple-go-api/internal/user/dto"

type VerifyEmptyName struct {
	candidate *dto.UserDto
}

func NewVerifyEmptyName(candidate *dto.UserDto) *VerifyEmptyName {
	return &VerifyEmptyName{
		candidate: candidate,
	}
}

func (spec VerifyEmptyName) IsSatisfiedBy(userDto dto.UserDto) (bool, string) {
	emptyString := ""

	if spec.candidate.Name == nil || spec.candidate.Name == &emptyString {
		return false, "O nome do usuário e obrigatório."
	}

	return true, ""
}
