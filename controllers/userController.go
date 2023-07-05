package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jouiwnl/go-crud/commons"
	"github.com/jouiwnl/go-crud/initializers"
	"github.com/jouiwnl/go-crud/models"
	"github.com/jouiwnl/go-crud/page"
)

func GetUSers(c *gin.Context) {
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))

	pageable := &page.Pageable{Offset: offset, Limit: limit}

	response := commons.FindAll[models.User](*pageable, initializers.Database)

	c.JSON(200, response)
}

func GetUser(c *gin.Context) {
	id := c.Param("id")

	dbContext := initializers.Database.Where("id = ?", id)

	response, err := commons.FindOne[models.User](dbContext)

	if err != nil {
		c.JSON(404, gin.H{
			"statusCode": 404,
			"message":    err.Error(),
		})
		return
	}

	c.JSON(200, response)
}
