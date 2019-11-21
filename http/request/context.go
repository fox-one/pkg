package request

import (
	"io/ioutil"
	"strings"

	"github.com/fox-one/pkg/text/localizer"
	"github.com/gin-gonic/gin"
)

const (
	localizerContextKey = "localizer_context_key"
)

func Body(c *gin.Context) (body []byte, err error) {
	if cb, ok := c.Get(gin.BodyBytesKey); ok {
		if cbb, ok := cb.([]byte); ok {
			body = cbb
		}
	}

	if body == nil {
		body, err = ioutil.ReadAll(c.Request.Body)
		if err == nil {
			c.Set(gin.BodyBytesKey, body)
		}
	}

	return
}

func Token(c *gin.Context) string {
	token := c.GetHeader("Authorization")
	return strings.TrimPrefix(token, "Bearer ")
}

func AcceptLanguage(c *gin.Context) string {
	return c.GetHeader("Accept-Language")
}

func WithLocalizer(l *localizer.Localizer) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(localizerContextKey, localizer.WithLanguage(l, AcceptLanguage(c)))
	}
}

func Localizer(c *gin.Context) *localizer.Localizer {
	return c.MustGet(localizerContextKey).(*localizer.Localizer)
}
