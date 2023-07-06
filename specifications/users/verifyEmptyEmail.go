package userSpecifications

import (
	"github.com/jouiwnl/go-crud/models"
)

type VerifyEmptyEmail struct {
}

func (spec VerifyEmptyEmail) IsSatisfiedBy(candidate models.User) (bool, string) {
	if candidate.Email == "" {
		return false, "O e-mail do usuário é obrigatório."
	}

	return true, ""
}
