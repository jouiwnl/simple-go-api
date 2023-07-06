package userSpecifications

import (
	"github.com/jouiwnl/go-crud/commons"
	"github.com/jouiwnl/go-crud/models"
	"gorm.io/gorm"
)

func GetSpecs(database *gorm.DB) []commons.Specification[models.User] {
	return []commons.Specification[models.User]{
		VerifyDuplicatedEmail{db: database},
		VerifyDuplicatedCpf{db: database},
		VerifyEmptyCpf{},
		VerifyEmptyEmail{},
		VerifyEmptyName{},
	}
}
