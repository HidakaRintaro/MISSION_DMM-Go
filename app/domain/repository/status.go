package repository

import (
	"context"
	"yatter-backend-go/app/domain/object"
)

type Status interface {
	// Create a status
	CreateStatus(ctx context.Context, accountId int64, status string) (int64, error)
	// Fetch status which has specified id
	FindById(ctx context.Context, id int64) (*object.Status, error)
	// Delete status which has specified id
	DeleteById(ctx context.Context, id int64) error
}
