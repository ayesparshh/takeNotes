package mocks

import (
	"time"

	"github.com/ayesparshh/internal/models"
)

type UserModel struct{}

func (m *UserModel) Insert(name, email, password string) error {
	switch email {
	case "dupe@example.com":
		return models.ErrDuplicateEmail
	default:
		return nil
	}
}
func (m *UserModel) Authenticate(email, password string) (int, error) {
	if email == "alice@example.com" && password == "pa$$word" {
		return 1, nil
	}
	return 0, models.ErrInvalidCredentials
}
func (m *UserModel) Exists(id int) (bool, error) {
	switch id {
	case 1:
		return true, nil
	default:
		return false, nil
	}
}

func (m *UserModel) Get(id int) (*models.User, error) {
	switch id {
	case 1:
		return &models.User{
			ID:             1,
			Name:           "Alice",
			Email:          "alice@example.com",
			HashedPassword: []byte("$2y$12$3.3/4.4/5.5/6.6/7.7/8.8/9.9/10.10/11.11/12.12"),
			Created:        time.Date(2022, 3, 17, 10, 15, 0, 0, time.UTC),
		}, nil
	default:
		return nil, models.ErrNoRecord
	}
}

func (m *UserModel) PasswordUpdate(id int, currentPassword, newPassword string) error {
	if id == 1 {
		if currentPassword != "pa$$word" {
			return models.ErrInvalidCredentials
		}
		return nil
	}
	return models.ErrNoRecord
}
