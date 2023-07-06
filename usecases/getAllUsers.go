package usecases

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jouiwnl/go-crud/initializers"
	"github.com/jouiwnl/go-crud/models"
	"github.com/jouiwnl/go-crud/page"
	"github.com/jouiwnl/go-crud/repository"
)

func GetAllUsers(c *gin.Context) {
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))

	pageable := page.Pageable{Offset: offset, Limit: limit}

	response := repository.FindAll[models.User](&pageable, initializers.Database)

	c.JSON(200, response)
}
