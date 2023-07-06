package repository

import (
	"errors"

	"github.com/jouiwnl/go-crud/page"
	"gorm.io/gorm"
)

func FindAll[T any](pageable *page.Pageable, dbContext *gorm.DB) page.Page[T] {
	var content []T
	dbContext.Offset(pageable.Offset).Limit(pageable.Limit).Find(&content)

	return page.Page[T]{
		Total:    len(content),
		Content:  content,
		Pageable: pageable,
	}
}

func FindOne[T any](dbContext *gorm.DB) (T, error) {
	var result T
	var err error

	dbResult := dbContext.Find(&result)

	if dbResult.RowsAffected <= 0 {
		err = errors.New("Não foi encontrado nenhum registro deste tipo com os filtros passados")
	}

	return result, err
}

func Save[T any](entity T, id any, dbContext *gorm.DB) T {
	if id == "" {
		dbContext.Create(&entity)
	} else {
		dbContext.Where("id = ?", id).Save(&entity)
	}

	return entity
}

func Exists[T any](dbContext *gorm.DB) bool {
	var content []T
	dbContext.Limit(1).First(&content)

	return len(content) > 0
}
