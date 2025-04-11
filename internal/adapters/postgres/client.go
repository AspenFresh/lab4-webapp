package postgres

import (
	"context"

	"github.com/jmoiron/sqlx"

	"github.com/AspenFresh/lab4-webapp/internal"
)

type Client struct {
	db *sqlx.DB
}

func NewClient(db *sqlx.DB) Client {
	return Client{db: db}
}

func (c Client) CreateUser(ctx context.Context, user internal.User) error {
	_, err := c.db.ExecContext(ctx, `
		INSERT INTO users (name, email, password, role_name) 
		VALUES ($1, $2, $3, $4)`,
		user.Name, user.Email, user.Password, user.Role.Name)
	return err
}

func (c Client) GetUserByEmail(ctx context.Context, email string) (internal.User, error) {
	var user internal.User
	var roleName string

	err := c.db.QueryRowxContext(ctx, `
		SELECT id, name, email, password, role_name 
		FROM users 
		WHERE email = $1`, email).
		Scan(&user.ID, &user.Name, &user.Email, &user.Password, &roleName)

	if err != nil {
		return internal.User{}, err
	}

	if roleName == "Admin" {
		user.Role = internal.Role{
			Name:        "Admin",
			Permissions: []string{"create", "read", "update", "delete"},
		}
	} else {
		user.Role = internal.Role{
			Name:        "User",
			Permissions: []string{"read"},
		}
	}

	return user, nil
}
