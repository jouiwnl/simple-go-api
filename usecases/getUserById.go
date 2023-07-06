package usecases

import (
	"github.com/gin-gonic/gin"
	"github.com/jouiwnl/go-crud/initializers"
	"github.com/jouiwnl/go-crud/models"
	"github.com/jouiwnl/go-crud/repository"
)

func GetUserById(c *gin.Context) {
	id := c.Param("id")

	dbContext := initializers.Database.Where("id = ?", id)

	response, err := repository.FindOne[models.User](dbContext)

	if err != nil {
		c.JSON(404, gin.H{
			"statusCode": 404,
			"message":    err.Error(),
		})
		return
	}

	c.JSON(200, response)
}
