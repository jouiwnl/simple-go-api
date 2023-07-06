package userSpecifications

import (
	"github.com/jouiwnl/go-crud/models"
)

type VerifyEmptyCpf struct {
}

func (spec VerifyEmptyCpf) IsSatisfiedBy(candidate models.User) (bool, string) {
	if candidate.Cpf == "" {
		return false, "O cpf do usuário e obrigatório."
	}

	return true, ""
}
