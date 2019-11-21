package store

import (
	"errors"

	"github.com/fox-one/pkg/store/db"
	"github.com/go-redis/redis"
)

var ErrNotFound = errors.New("store: nil")

func IsErrNotFound(err error) bool {
	for err != nil {
		if err == ErrNotFound || err == redis.Nil || db.IsErrorNotFound(err) {
			return true
		}

		err = errors.Unwrap(err)
	}

	return false
}

func IsErrOptimisticLock(err error) bool {
	return errors.Is(err, db.ErrOptimisticLock)
}
