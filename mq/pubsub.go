package mq

import (
	"context"
	"time"
)

type PublishOption struct {
	GroupID   string
	TraceID   string
	VisibleAt time.Time
	ExpiredAt time.Time
}

type Pub interface {
	Publish(ctx context.Context, msg string, opt *PublishOption) error
}

type Callback interface {
	// Finish finish the message
	Finish(ctx context.Context) error
	// Delay make the message visible again after interval
	Delay(ctx context.Context, duration time.Duration) error
}

type ReceiveOption struct {
	VisibilityTimeout time.Duration
}

type Sub interface {
	Receive(ctx context.Context, opt *ReceiveOption) (string, Callback, error)
}

type PubSub interface {
	Pub
	Sub
}
