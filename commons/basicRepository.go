package commons

import (
	"errors"

	"github.com/jouiwnl/go-crud/page"
	"gorm.io/gorm"
)

func FindAll[T any](pageable page.Pageable, dbContext *gorm.DB) page.Page[T] {
	var content []T
	dbContext.Offset(pageable.Offset).Limit(pageable.Limit).Find(&content)

	return page.Page[T]{
		Total:    len(content),
		Content:  content,
		Pageable: &pageable,
	}
}

func FindOne[T any](dbContext *gorm.DB) (T, error) {
	var result T
	var err error

	dbResult := dbContext.Find(&result)

	if dbResult.RowsAffected <= 0 {
		err = errors.New("Não foi possível encontrar um registro deste tipo com os filtros passados")
	}

	return result, err
}

/*
func Save[T utils.IEntity](entity T, dbContext *gorm.DB) (T, error) {
	var err error

	id := entity.GetId();

	dbContext.

	dbResult := dbContext.Find(&result)

	if dbResult.RowsAffected <= 0 {
		err = errors.New("Não foi possível encontrar um registro deste tipo com os filtros passados")
	}

	return result, err
}*/
