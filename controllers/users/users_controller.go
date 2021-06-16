package users

import (
	"github.com/burnout09/bookstore-users-api/domain/users"
	"github.com/burnout09/bookstore-users-api/services"
	"github.com/burnout09/bookstore-users-api/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateUser(c *gin.Context) {
	var user users.User

	//bytes, err := ioutil.ReadAll(c.Request.Body)
	//if err != nil {
	//	//TODO: Handle errors
	//	fmt.Println(err)
	//	return
	//}
	//
	//if err := json.Unmarshal(bytes, &user); err != nil {
	//	//TODO: Handle json errors
	//	fmt.Println(err)
	//	return
	//}

	//ShouldBindJSON equivalente de usar ioutil.ReadAll y json.Unmarshal
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("Invalid json body")
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

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}
