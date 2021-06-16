package users

import (
	"fmt"
	"github.com/burnout09/bookstore-users-api/utils/date_utils"
	"github.com/burnout09/bookstore-users-api/utils/errors"
)

var (
	usersBD = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr {
	result := usersBD[user.Id]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.Id))
	}
	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated

	return nil
}

func (user *User) Save() *errors.RestErr {
	current := usersBD[user.Id]
	if current != nil {
		if current.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprintf("email %s already registered", user.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("user %d already exists", user.Id))
	}

	user.DateCreated = date_utils.GetNowString()

	usersBD[user.Id] = user

	return nil
}
