package users

import (
	"fmt"
	"github.com/burnout09/bookstore-users-api/domain/users"
	"github.com/burnout09/bookstore-users-api/services"
	"github.com/burnout09/bookstore-users-api/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func getUserId(userIdParam string) (int64, *errors.RestErr) {
	userId, err := strconv.ParseInt(userIdParam, 10, 64)
	if err != nil {
		return 0, errors.NewBadRequestError("invalid user id, user id should be a number")
	}
	return userId, nil
}

func Create(c *gin.Context) {
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

	var user users.User
	//ShouldBindJSON equivalente de usar ioutil.ReadAll y json.Unmarshal
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("Invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.UsersService.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result.Marshall(c.GetHeader("X-Public") == "true"))
}

func Get(c *gin.Context) {
	userId, err := getUserId(c.Param("user_id"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	user, getErr := services.UsersService.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	c.JSON(http.StatusOK, user.Marshall(c.GetHeader("X-Public") == "true"))
}

func Update(c *gin.Context) {
	userId, err := getUserId(c.Param("user_id"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("Invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	user.Id = userId

	isPartial := c.Request.Method == http.MethodPatch

	result, err := services.UsersService.UpdateUser(isPartial, user)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, result.Marshall(c.GetHeader("X-Public") == "true"))
}

func Delete(c *gin.Context) {
	userId, err := getUserId(c.Param("user_id"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	if err := services.UsersService.DeleteUser(userId); err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

func Search(c *gin.Context) {
	status := c.Query("status")
	users, err := services.UsersService.SearchUser(status)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, users.Marshall(c.GetHeader("X-Public") == "true"))
}

func Login(c *gin.Context) {
	var request users.LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		restError := errors.NewBadRequestError("invalid json body")
		c.JSON(restError.Status, restError)
		return
	}
	user, err := services.UsersService.LoginUser(request)
	fmt.Println("LOGIN")
	fmt.Println(err)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	//c.JSON(http.StatusOK, user)
	c.JSON(http.StatusOK, user.Marshall(c.GetHeader("X-Public") == "true"))
}
