package auth

import (
	"context"
)

type Auth interface {
	Verify(ctx context.Context, password string) (bool, error)
}
