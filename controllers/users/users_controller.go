package users

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/parsaakbari1209/go-mongo-CRUD-web/domain"
	"github.com/parsaakbari1209/go-mongo-CRUD-web/services"
	"github.com/parsaakbari1209/go-mongo-CRUD-web/utils"
)

func CreateUser(c *gin.Context) {
	var newUser domain.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		restErr := utils.BadRequest("Invalid json.")
		c.JSON(restErr.Status, restErr)
		return
	}
	user, restErr := services.CreateUser(&newUser)
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}
	c.JSON(http.StatusCreated, user)
}

func FindUser(c *gin.Context) {
	userEmail := c.Query("email")
	if userEmail == "" {
		restErr := utils.BadRequest("No email.")
		c.JSON(restErr.Status, restErr)
		return
	}
	fmt.Println(userEmail)
	user, restErr := services.FindUser(userEmail)
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}
	c.JSON(http.StatusOK, user)
}