package services

import (
	"github.com/burnout09/bookstore-users-api/domain/users"
	"github.com/burnout09/bookstore-users-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	return &user, nil
}
