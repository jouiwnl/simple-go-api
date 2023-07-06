package userSpecifications

import (
	"fmt"

	"github.com/jouiwnl/go-crud/models"
	"github.com/jouiwnl/go-crud/repository"
	"gorm.io/gorm"
)

type VerifyDuplicatedCpf struct {
	db *gorm.DB
}

func (spec VerifyDuplicatedCpf) IsSatisfiedBy(candidate models.User) (bool, string) {
	query := spec.db.Where("cpf like ?", candidate.Cpf)

	if candidate.ID != "" {
		query.Where("id <> ?", candidate.ID)
	}

	exists := repository.Exists[models.User](query)

	if exists {
		return false, fmt.Sprintf("Já existe um usuário cadastrado com o cpf %s.", candidate.Cpf)
	}

	return true, ""
}
