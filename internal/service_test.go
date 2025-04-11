package internal

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

type MockStorage struct {
	users []User
}

func (m *MockStorage) CreateUser(ctx context.Context, user User) error {
	m.users = append(m.users, user)
	return nil
}

func (m *MockStorage) GetUserByEmail(ctx context.Context, email string) (User, error) {
	for _, user := range m.users {
		if user.Email == email {
			return user, nil
		}
	}
	return User{}, errors.New("user not found")
}

func TestCreateUser(t *testing.T) {
	mock := &MockStorage{}
	service := NewUserService(mock)

	user := User{
		Name:     "Maryna",
		Email:    "maryna@example.com",
		Password: "changeme",
		Role: Role{
			Name:        "Admin",
			Permissions: []string{"create", "read", "update", "delete"},
		},
	}

	createdUser, err := service.CreateUser(context.Background(), user)
	require.NoError(t, err)
	require.Equal(t, 1, len(mock.users))
	require.Equal(t, "Maryna", createdUser.Name)
}

func TestGetUserByEmail_Success(t *testing.T) {
	mock := &MockStorage{
		users: []User{
			{
				Name:  "Ivan",
				Email: "ivan@example.com",
				Role: Role{
					Name:        "User",
					Permissions: []string{"read"},
				},
			},
		},
	}
	service := NewUserService(mock)

	user, err := service.GetUserByEmail(context.Background(), "ivan@example.com")
	require.NoError(t, err)
	require.Equal(t, "Ivan", user.Name)
	require.True(t, user.HasAccess("read"))
	require.False(t, user.HasAccess("delete"))
}

func TestGetUserByEmail_NotFound(t *testing.T) {
	mock := &MockStorage{}
	service := NewUserService(mock)

	_, err := service.GetUserByEmail(context.Background(), "notfound@example.com")
	require.Error(t, err)
}
