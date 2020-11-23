package services

import (
	"github.com/kaitus/bookstore_users-api-golang/domain/users"
	"github.com/kaitus/bookstore_users-api-golang/utils/errors"
)

func GetUser(userId int64) (*users.User, *errors.RestError) {
	result := &users.User{ Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

func CreateUser(user users.User) (*users.User, *errors.RestError) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func UpdateUser(isPatial bool, user users.User) (*users.User, *errors.RestError) {
	current, err := GetUser(user.Id)
	if err != nil {
		return nil, err
	}

	if err := user.Validate(); err != nil {
		return nil, err
	}

	if isPatial {
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}
		if user.LastName != "" {
			current.LastName = user.LastName
		}
		if user.Email != "" {
			current.Email = user.Email
		}
	} else {
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
	}


	if err := current.Update(); err != nil {
		return nil, err
	}
	return current, nil
}

func DeleteUser(userId int64) *errors.RestError {
	user := &users.User{Id: userId}
	return user.Delete()
}