package repository

import (
	"context"

	"yatter-backend-go/app/domain/object"
)

type Account interface {
	// Fetch account which has specified username
	FindByUsername(ctx context.Context, username string) (*object.Account, error)
	// Create an account
	CreateAccount(ctx context.Context, username string, passwordHash string) (int64, error)
	// Fetch account which has specified id
	FindById(ctx context.Context, id int64) (*object.Account, error)
}
