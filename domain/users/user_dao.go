package users

import (
	"fmt"
	"github.com/kaitus/bookstore_users-api-golang/utils/date_utils"
	"github.com/kaitus/bookstore_users-api-golang/utils/errors"
	"github.com/kaitus/bookstore_users-api-golang/datasources/mysql/users_db"
)

var (
	usersDB = make(map[int64] *User)
)

func (user *User) Get() *errors.RestError {
	if err := users_db.Client.Ping(); err != nil {
		panic(err)
	}
	result := usersDB[user.Id]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found ", user.Id))
	}
	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated
	return nil
}

func (user *User) Save() *errors.RestError {
	current := usersDB[user.Id]
	if current != nil {
		if current.Email == user.Email {
			return errors.NewBadRequest(fmt.Sprintf("email %s already registered", user.Email))
		}
		return errors.NewBadRequest(fmt.Sprintf("user %d already exist", user.Id))
	}
	user.DateCreated = date_utils.GetNowString()
	usersDB[user.Id] =user
	return nil
}
