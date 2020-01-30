package domain_test

import (
	"github.com/berlin-ab/golang-scratch-pad/dependencyInversionExample/domain"
	"github.com/berlin-ab/golang-scratch-pad/dependencyInversionExample/implementation"
	"testing"
)

func TestExample(t *testing.T) {
	t.Run("async stuff", func(t *testing.T) {
		userRepository := implementation.NewUserRepository()
		users, _ := domain.UpperCaseUsers(userRepository)

		if users[0].Name != "JOE" {
			t.Errorf("expected first user to be JOE, got %v", users[0].Name)
		}

		if users[1].Name != "JACK" {
			t.Errorf("expected second user to be JACK, got %v", users[1].Name)
		}
	})
}
