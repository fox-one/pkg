package param

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestBindingParams(t *testing.T) {
	{
		r, _ := http.NewRequest("GET", "https://api.fox.one?symbol=BOX&amount=0.1&hide=1&foo=bar", nil)

		var params struct {
			Symbol string          `json:"symbol,omitempty"`
			Amount decimal.Decimal `json:"amount,omitempty"`
			Hide   bool            `json:"hide,omitempty"`
		}

		if err := Binding(r, &params); assert.Nil(t, err) {
			assert.Equal(t, "BOX", params.Symbol)
			assert.Equal(t, "0.1", params.Amount.String())
			assert.True(t, params.Hide)
		}
	}

	{
		// binding body only
		body := strings.NewReader(`{"body":"body value"}`)
		r, _ := http.NewRequest("POST", "https://example.com?symbol=BOX&amount=0.1&hide=1&foo=bar", body)

		var params struct {
			Symbol string          `json:"symbol,omitempty"`
			Amount decimal.Decimal `json:"amount,omitempty"`
			Hide   bool            `json:"hide,omitempty"`
			Body   string          `json:"body"`
		}

		if err := Binding(r, &params); assert.Nil(t, err) {
			assert.Empty(t, params.Symbol)
			assert.Empty(t, params.Amount)
			assert.Empty(t, params.Hide)
			assert.Equal(t, "body value", params.Body)
		}
	}
}

func TestBindingBoth(t *testing.T) {
	var result struct {
		// url params
		ID int `json:"id"`
		// url query params
		Foo string `json:"foo"`
		// body
		Body string `json:"body"`
	}
	router := chi.NewRouter()
	router.Post("/{id}", func(rw http.ResponseWriter, r *http.Request) {
		assert.Nil(t, BindingBoth(r, &result))
	})

	reader := strings.NewReader(`{"body":"body value"}`)
	req := httptest.NewRequest("POST", "http://example.com/10?foo=bar", reader)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 10, result.ID)
	assert.Equal(t, "bar", result.Foo)
	assert.Equal(t, "body value", result.Body)
}
