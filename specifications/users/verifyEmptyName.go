package userSpecifications

import (
	"github.com/jouiwnl/go-crud/models"
)

type VerifyEmptyName struct {
}

func (spec VerifyEmptyName) IsSatisfiedBy(candidate models.User) (bool, string) {
	if candidate.Name == "" {
		return false, "O nome do usuário e obrigatório."
	}

	return true, ""
}
