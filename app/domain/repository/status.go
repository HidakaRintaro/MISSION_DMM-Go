package repository

import "context"

type Status interface {
	// Create a status
	CreateStatus(ctx context.Context, status string) (int64, error)
}
