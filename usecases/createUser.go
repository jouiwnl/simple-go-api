package usecases

import (
	"github.com/gin-gonic/gin"
	specificationHelper "github.com/jouiwnl/go-crud/helpers"
	"github.com/jouiwnl/go-crud/initializers"
	"github.com/jouiwnl/go-crud/models"
	"github.com/jouiwnl/go-crud/repository"
	userSpecifications "github.com/jouiwnl/go-crud/specifications/users"
)

func CreateUser(c *gin.Context) {
	user := models.User{}
	c.BindJSON(&user)

	database := initializers.Database

	errors := specificationHelper.RunSpecs[models.User](user, userSpecifications.GetSpecs(database))

	if len(errors) > 0 {
		c.JSON(422, gin.H{
			"message": "Erro ao processar entidade",
			"errors":  errors,
		})
		return
	}

	c.JSON(200, repository.Save[models.User](user, "", database))
}
