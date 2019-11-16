package request

import (
	"io/ioutil"
	"strings"

	"github.com/gin-gonic/gin"
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
