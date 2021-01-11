package mysql_utils

import (
	"github.com/go-sql-driver/mysql"
	"github.com/kaitus/bookstore_users-api-golang/utils/errors"
	"strings"
)

const (
	ErrorNoRows = "no rows in result set"
)

func ParseError(err error) *errors.RestError {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), ErrorNoRows) {
			return errors.NewNotFoundError("no record matching given id")
		}
		return errors.NewInternalServerError("error parsing database response")
	}

	switch sqlErr.Number {
	case 1062:
		return errors.NewBadRequest(sqlErr.Message)
	}
	return errors.NewInternalServerError("error processing request")
}