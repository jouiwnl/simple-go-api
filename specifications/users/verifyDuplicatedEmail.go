package userSpecifications

import (
	"fmt"

	"github.com/jouiwnl/go-crud/models"
	"github.com/jouiwnl/go-crud/repository"
	"gorm.io/gorm"
)

type VerifyDuplicatedEmail struct {
	db *gorm.DB
}

func (spec VerifyDuplicatedEmail) IsSatisfiedBy(candidate models.User) (bool, string) {
	query := spec.db.Where("email like ?", candidate.Email)

	if candidate.ID != "" {
		query.Where("id <> ?", candidate.ID)
	}

	exists := repository.Exists[models.User](query)

	if exists {
		return false, fmt.Sprintf("Já existe um usuário cadastrado com o e-mail %s.", candidate.Email)
	}

	return true, ""
}
