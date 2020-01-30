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

func UpperCaseUsers(userRepository UserRepository) ([]User, error) {
	success := make(chan []User, 1)
	error := make(chan error, 1)

	userRepository.GetUsers(success, error)

	select {
	case result := <-success:
		return transformUserNames(result), nil
	case <-error:
		return nil, buildFailure()
	}
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
