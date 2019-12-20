package scannerapi

import (
	"context"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	_ = iota
	contextKeyApp
)

type App struct {
	AppID      string
	PrivateKey *rsa.PrivateKey
}

func (app App) Auth(r *http.Request) error {
	url := r.URL.String()
	idx := strings.Index(url, r.URL.Path)
	uri := url[idx:]

	var body []byte
	if r.GetBody != nil {
		if rc, _ := r.GetBody(); rc != nil {
			defer rc.Close()
			body, _ = ioutil.ReadAll(rc)
		}
	}

	h := sha256.New()
	h.Write(append([]byte(r.Method+uri), body...))
	digest := h.Sum(nil)

	bts, err := rsa.SignPKCS1v15(rand.Reader, app.PrivateKey, crypto.SHA256, digest)
	if err == nil {
		r.Header.Add("Authorization", base64.StdEncoding.EncodeToString(bts))
	}
	return err
}

func WithApp(ctx context.Context, app *App) context.Context {
	return context.WithValue(ctx, contextKeyApp, app)
}
