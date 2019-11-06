package logger

import (
	"fmt"
	"net/http"
	"time"

	"github.com/fox-one/pkg/uuid"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.GetHeader("X-Request-ID")
		if id == "" {
			id = uuid.New()
		}

		ctx := c.Request.Context()
		log := FromContext(ctx).WithField("request-id", id)
		ctx = WithContext(ctx, log)
		c.Request = c.Request.WithContext(ctx)

		start := time.Now()
		c.Next()
		end := time.Now()
		status := c.Writer.Status()
		method := c.Request.Method
		uri := c.Request.URL.String()

		content := fmt.Sprintf("[%d] %-4s %s", status, method, uri)

		log = log.WithFields(logrus.Fields{
			"ts": start.Format(time.RFC3339),
			"lt": end.Sub(start),
			"ip": c.ClientIP(),
			"ua": c.Request.UserAgent(),
		})

		switch {
		case status >= http.StatusOK && status < 300:
			log.Info(content)
		case status == http.StatusNotFound:
			log.Warn(content)
		default:
			log.Error(content)
		}
	}
}
