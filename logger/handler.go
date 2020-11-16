package logger

import (
	"fmt"
	"net/http"
	"time"

	"github.com/fox-one/pkg/uuid"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const (
	RequestIdHeaderKey = "X-Request-Id"
	RequestIdLogKey    = "request-id"
	ExposeHeaderKey    = "Access-Control-Expose-Headers"
)

func Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.GetHeader(RequestIdHeaderKey)
		if id == "" {
			id = uuid.New()
		}

		ctx := c.Request.Context()
		log := FromContext(ctx).WithField(RequestIdLogKey, id)
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

// Middleware provides logging middleware.
func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.Header.Get(RequestIdHeaderKey)
		if id == "" {
			id = uuid.New()
		}
		ctx := r.Context()
		log := FromContext(ctx).WithField(RequestIdLogKey, id)
		ctx = WithContext(ctx, log)
		start := time.Now()
		next.ServeHTTP(w, r.WithContext(ctx))
		end := time.Now()
		log.WithFields(logrus.Fields{
			"method":  r.Method,
			"request": r.RequestURI,
			"remote":  r.RemoteAddr,
			"latency": end.Sub(start),
			"time":    end.Format(time.RFC3339),
		}).Debug()
	})
}

func WithRequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.Header.Get(RequestIdHeaderKey)
		if id == "" {
			id = uuid.New()
		}

		w.Header().Set(RequestIdHeaderKey, id)
		w.Header().Set(ExposeHeaderKey, RequestIdHeaderKey)

		ctx := r.Context()
		log := FromContext(ctx).WithField(RequestIdLogKey, id)
		ctx = WithContext(ctx, log)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
