package implementation

import (
	domain "github.com/berlin-ab/golang-scratch-pad/dependencyInversionExample/domain"
)

type userRepositoryData struct {
}

func NewUserRepository() *userRepositoryData {
	return &userRepositoryData{}
}

func (client *userRepositoryData) GetUsers(successChannel chan<- []domain.User, errorChannel chan<- error) {
	go func() {
		successChannel <- []domain.User{
			{"Joe"},
			{"Jack"},
		}
	}()
}

