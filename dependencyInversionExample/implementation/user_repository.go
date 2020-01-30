package implementation

import domain "github.com/berlin-ab/golang-scratch-pad/dependencyInversionExample/domain"

type userRepositoryData struct {
}

func NewUserRepository() *userRepositoryData {
	return &userRepositoryData{}
}

func (client *userRepositoryData) GetUsers(success chan<- []domain.User, error chan<- error) {
	go func() {
		success <- []domain.User{
			{"Joe"},
			{"Jack"},
		}
	}()
}

