package property

import (
	"context"
	"testing"
	"time"

	"github.com/fox-one/pkg/store/database"
	"github.com/stretchr/testify/assert"
)

func TestPropertyStore(t *testing.T) {
	db, err := database.Open(database.SqliteInMemory())
	if err != nil {
		t.Fatal(err)
	}

	if err := database.Migrate(db); err != nil {
		t.Fatal(err)
	}

	s := New(db)
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
}
