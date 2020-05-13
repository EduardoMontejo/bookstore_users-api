package users

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/EduardoMontejo/bookstore_users-api/domain/users"
	"github.com/EduardoMontejo/bookstore_users-api/services"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user users.User
	fmt.Println(user)
	bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		//TODO: Handle error
		return
	}

	if err := json.Unmarshal(bytes, &user); err != nil {
		fmt.Println(err.Error())
		//TODO: Handle json error
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		//TODO handle user creation error
		return
	}
	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}
