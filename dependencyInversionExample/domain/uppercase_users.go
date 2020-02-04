package domain

import (
	"errors"
	"fmt"
	"strings"
)

type User struct {
	Name string
}

type UserRepository interface {
	GetUsers(a chan<- []User, b chan<- error)
}

func UpperCaseUsers(userRepositories []UserRepository) ([]User, error) {
	success := make(chan []User, len(userRepositories))
	errorChannel := make(chan error, len(userRepositories))

	for _, userRepository := range userRepositories {
		userRepository.GetUsers(success, errorChannel)
	}

	var users []User
	var errs []error

	for range userRepositories {
		select {
		case result := <-success:
			users = append(users, result...)
		case err := <-errorChannel:
			errs = append(errs, err)
		}
	}

	if len(errs) > 0 {
		return nil, buildFailure()
	}

	return transformUserNames(users), nil
}

func transformUserNames(users []User) []User {
	var resultUsers []User

	for _, user := range users {
		user.Name = strings.ToUpper(user.Name)
		resultUsers = append(resultUsers, user)
	}

	return resultUsers
}

func buildFailure() error {
	fmt.Print("fail")
	return errors.New("failed")
}
