package property

import (
	"context"

	"github.com/yiplee/go-cache"
)

func Cache(store Store) Store {
	return &cacheProperties{
		Store: store,
		cache: cache.New[string, Value](),
	}
}

type cacheProperties struct {
	Store
	cache *cache.Cache[string, Value]
}

func (s *cacheProperties) Get(ctx context.Context, key string) (Value, error) {
	if v, ok := s.cache.Get(key); ok {
		return v, nil
	}

	v, err := s.Store.Get(ctx, key)
	if err == nil {
		s.cache.Set(key, v)
	}

	return v, nil
}

func (s *cacheProperties) Save(ctx context.Context, key string, value interface{}) error {
	if err := s.Store.Save(ctx, key, value); err != nil {
		return err
	}

	s.cache.Set(key, Parse(value))
	return nil
}

func (s *cacheProperties) Expire(ctx context.Context, key string) error {
	if err := s.Store.Expire(ctx, key); err != nil {
		return err
	}

	s.cache.Delete(key)
	return nil
}
