package propertystore

import (
	"context"
	"testing"
	"time"

	dbcommon "github.com/fox-one/pkg/db"
	"github.com/fox-one/pkg/store/db"
	"github.com/stretchr/testify/assert"
)

func TestPropertyStore(t *testing.T) {
	dbs, err := db.Open(dbcommon.SqliteInMemory())
	if err != nil {
		t.Fatal(err)
	}

	if err := db.Migrate(dbs); err != nil {
		t.Fatal(err)
	}

	s := New(dbs)
	ctx := context.Background()

	values := map[string]interface{}{
		"number": 1234,
		"string": "hahahha",
		"time":   time.Now(),
	}

	t.Run("save property", func(t *testing.T) {
		for k, v := range values {
			err := s.Save(ctx, k, v)
			assert.Nil(t, err, "save %s should be ok", k)
		}
	})

	t.Run("get property", func(t *testing.T) {
		for k := range values {
			v, err := s.Get(ctx, k)
			if assert.Nil(t, err) {
				t.Log(k, v.String())
				assert.NotEmpty(t, v.String())
			}
		}
	})

	t.Run("list values", func(t *testing.T) {
		list, err := s.List(ctx)
		if assert.Nil(t, err) {
			assert.Len(t, list, 3)

			for k, v := range list {
				t.Log(k, v.String())
			}
		}
	})

	t.Run("expire property", func(t *testing.T) {
		for k := range values {
			err := s.Expire(ctx, k)
			assert.Nil(t, err)
		}
	})

	t.Run("list values", func(t *testing.T) {
		list, err := s.List(ctx)
		assert.Nil(t, err)
		assert.Len(t, list, 0)
	})
}
