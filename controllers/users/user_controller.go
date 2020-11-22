package users

import (
	"github.com/gin-gonic/gin"
	"github.com/kaitus/bookstore_users-api-golang/domain/users"
	"github.com/kaitus/bookstore_users-api-golang/services"
	"github.com/kaitus/bookstore_users-api-golang/utils/errors"
	"net/http"
	"strconv"
)

func GetUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequest("User id should be a number")
		c.JSON(err.Status, err)
		return
	}
	user, getErr := services.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, user)
}

func CreateUser(c *gin.Context)  {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequest("Invalid Json Body")
		c.JSON(restErr.Status, restErr)
		return
	}
	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}
