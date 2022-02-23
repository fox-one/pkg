package param

import (
	"encoding/json"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/go-chi/chi"
	"github.com/spf13/cast"
	"github.com/twitchtv/twirp"
)

// String read param in path or query
func String(r *http.Request, key string) string {
	v := chi.URLParam(r, key)
	if v == "" {
		v = r.URL.Query().Get(key)
	}

	return v
}

func Int(r *http.Request, key string) int {
	return cast.ToInt(String(r, key))
}

func Int64(r *http.Request, key string) int64 {
	return cast.ToInt64(String(r, key))
}

// Binding decode request params(url params or body) to struct with json tag
func Binding(r *http.Request, v interface{}) error {
	return binding(r, v, false)
}

// BindingBoth decode request params (url params and body) to struct with json tag
// NOTE: if the method is not PATCH/POST/PUT, the request body will be ignored.
func BindingBoth(r *http.Request, v interface{}) error {
	return binding(r, v, true)
}

func binding(r *http.Request, v interface{}, both bool) error {
	var err error

	bodyBound := false
	switch r.Method {
	case http.MethodPatch, http.MethodPost, http.MethodPut:
		bodyBound = true
		err = bindingBody(r, v)
	}

	if !bodyBound || both {
		if err == nil {
			err = bindingParams(r, v)
		}
	}

	if err == nil {
		if _, verr := govalidator.ValidateStruct(v); verr != nil {
			err = twirp.NewError(twirp.InvalidArgument, verr.Error())
		}
	}

	return err
}

func bindingBody(r *http.Request, v interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		return twirp.NewError(twirp.Malformed, "can't decode request body")
	}

	return nil
}

func bindingParams(r *http.Request, v interface{}) error {
	values := r.URL.Query()

	if ctx := chi.RouteContext(r.Context()); ctx != nil {
		params := ctx.URLParams
		for idx := range params.Keys {
			key, value := params.Keys[idx], params.Values[idx]
			values.Set(key, value)
		}
	}

	if len(values) > 0 {
		if err := globalDecoder.Decode(v, values); err != nil {
			return twirp.NewError(twirp.InvalidArgument, err.Error())
		}
	}

	return nil
}
