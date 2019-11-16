package property

import "context"

type Store interface {
	Get(ctx context.Context, key string) (Value, error)
	Save(ctx context.Context, key string, value interface{}) error
	List(ctx context.Context) (map[string]Value, error)
}
