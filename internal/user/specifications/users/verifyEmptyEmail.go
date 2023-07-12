package userSpecifications

import "github.com/jouiwnl/simple-go-api/internal/user/dto"

type VerifyEmptyEmail struct {
	candidate *dto.UserDto
}

func NewVerifyEmptyEmail(candidate *dto.UserDto) *VerifyEmptyEmail {
	return &VerifyEmptyEmail{
		candidate: candidate,
	}
}

func (spec VerifyEmptyEmail) IsSatisfiedBy(userDto dto.UserDto) (bool, string) {
	emptyString := ""

	if spec.candidate.Email == nil || spec.candidate.Email == &emptyString {
		return false, "O e-mail do usuário e obrigatório."
	}

	return true, ""
}
