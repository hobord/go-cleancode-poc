package auth

import (
	"context"
)

// Auth is the interface of the authentication
type Auth interface {
	Verify(ctx context.Context, userID, password string) (bool, error)
	SetPassword(ctx context.Context, userID, password string) error
}
