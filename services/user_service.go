package services

import (
	"github.com/burnout09/bookstore-users-api/domain/users"
	"github.com/burnout09/bookstore-users-api/utils/errors"
	"strings"
)

func GetUser(userId int64) (*users.User, *errors.RestErr) {
	result := users.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return &result, nil
}

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

func UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestErr) {
	current, err := GetUser(user.Id)
	if err != nil {
		return nil, err
	}

	if isPartial {
		if len(strings.TrimSpace(user.FirstName)) > 0 {
			current.FirstName = user.FirstName
		}
		if len(strings.TrimSpace(user.LastName)) > 0 {
			current.LastName = user.LastName
		}
		if len(strings.TrimSpace(user.Email)) > 0 {
			current.Email = user.Email
		}
	} else {
		if err := user.Validate(); err != nil {
			return nil, err
		}
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
	}

	if err := current.Update(); err != nil {
		return nil, err
	}
	return current, nil
}
