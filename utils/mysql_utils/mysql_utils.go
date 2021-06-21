package mysql_utils

import (
	"github.com/VividCortex/mysqlerr"
	"github.com/burnout09/bookstore-users-api/utils/errors"
	"github.com/go-sql-driver/mysql"
	"strings"
)

const (
	ErrorNoRows = "no rows in result set"
)

func ParseError(err error) *errors.RestErr {
	sqlError, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), ErrorNoRows) {
			return errors.NewNotFoundError("no record matching given id")
		}
		return errors.NewInternalServerError("error parsing database response")
	}

	switch sqlError.Number {
	case mysqlerr.ER_DUP_ENTRY:
		return errors.NewBadRequestError("invalid data, duplicated key")
	default:
		return errors.NewInternalServerError("error processing request")
	}
}
