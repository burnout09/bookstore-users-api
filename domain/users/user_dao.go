package users

import (
	"fmt"
	"github.com/burnout09/bookstore-users-api/datasources/mysql/users_db"
	"github.com/burnout09/bookstore-users-api/utils/date_utils"
	"github.com/burnout09/bookstore-users-api/utils/errors"
	"strings"
)

const (
	indexUniqueEmail = "users_email_uindex"
	errorNoRows      = "no rows in result set"
	queryInsertUser  = "INSERT INTO users(first_name,last_name,email,date_created) VALUES(?,?,?,?);"
	queryGetUser     = "SELECT id,first_name,last_name,email,date_created FROM users WHERE id=?;"
)

var (
	usersBD = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Id)
	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); err != nil {
		fmt.Println(err)
		if strings.Contains(err.Error(), errorNoRows) {
			return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.Id))
		}
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to get user %d: %s", user.Id, err.Error()))
	}

	return nil
}

func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	user.DateCreated = date_utils.GetNowString()

	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if err != nil {
		if strings.Contains(err.Error(), indexUniqueEmail) {
			return errors.NewBadRequestError(fmt.Sprintf("email %s already exists", user.Email))
		}
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to save user :%s", err.Error()))
	}

	userId, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to save user :%s", err.Error()))
	}

	user.Id = userId

	//current := usersBD[user.Id]
	//if current != nil {
	//	if current.Email == user.Email {
	//		return errors.NewBadRequestError(fmt.Sprintf("email %s already registered", user.Email))
	//	}
	//	return errors.NewBadRequestError(fmt.Sprintf("user %d already exists", user.Id))
	//}
	//

	//
	//usersBD[user.Id] = user

	return nil
}
