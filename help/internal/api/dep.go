package api

import "context"

type AnekdotProvider interface {
	GetAnekdot(ctx context.Context, category string) (string, error)
}
