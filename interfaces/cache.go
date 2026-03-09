package interfaces

import (
	"context"

	"github.com/saianand32/gocachex/internal/options"
)

type Cache[K comparable, V any] interface {
	Get(ctx context.Context, key K) (V, bool)
	Set(ctx context.Context, key K, value V, opts ...options.Option) error
	Delete(ctx context.Context, key K) bool
	Clear(ctx context.Context)
	Len() int
	GetType() string
	Close()
}
