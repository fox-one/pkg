package sqs

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/fox-one/pkg/mq"
	jsoniter "github.com/json-iterator/go"
)

type queue struct {
	client   *sqs.SQS
	queueURL *string
	fifo     bool
}

func New(s *session.Session, queueURL string) mq.PubSub {
	return &queue{
		client:   sqs.New(s),
		queueURL: aws.String(queueURL),
		fifo:     isFifoQueue(queueURL),
	}
}

type message struct {
	Content   string    `json:"content,omitempty"`
	VisibleAt time.Time `json:"visible_at,omitempty"`
	ExpiredAt time.Time `json:"expired_at,omitempty"`
}

func (msg *message) IsExpired() bool {
	return !msg.ExpiredAt.IsZero() && msg.ExpiredAt.Before(time.Now())
}

func (q *queue) Publish(ctx context.Context, content string, opt *mq.PublishOption) error {
	msg := message{
		Content:   content,
		VisibleAt: time.Now(),
	}

	if opt != nil {
		if !opt.ExpiredAt.IsZero() {
			msg.ExpiredAt = opt.ExpiredAt
		}

		if !q.fifo {
			if opt.VisibleAt.After(msg.VisibleAt) {
				msg.VisibleAt = opt.VisibleAt
			}
		}
	}

	body, _ := jsoniter.MarshalToString(msg)
	input := &sqs.SendMessageInput{
		MessageBody: aws.String(body),
		QueueUrl:    q.queueURL,
	}

	if d := time.Until(msg.VisibleAt); d > 0 {
		// sqs 普通队列最长 delay 15min
		if d > 15*time.Minute {
			d = 15 * time.Minute
		}

		input.DelaySeconds = aws.Int64(d.Milliseconds() / 1000)
	}

	if q.fifo && opt != nil {
		input.MessageGroupId = aws.String(opt.GroupID)

		if opt.TraceID != "" {
			input.MessageDeduplicationId = aws.String(opt.TraceID)
		}
	}

	_, err := q.client.SendMessageWithContext(ctx, input)
	return err
}

func (q *queue) Receive(ctx context.Context, opt *mq.ReceiveOption) (string, mq.Callback, error) {
	input := &sqs.ReceiveMessageInput{
		QueueUrl:          q.queueURL,
		WaitTimeSeconds:   aws.Int64(10),
		VisibilityTimeout: aws.Int64(10),
	}

	if opt != nil {
		if timeout := int64(opt.VisibilityTimeout.Seconds()); timeout > 0 {
			input.VisibilityTimeout = aws.Int64(timeout)
		}
	}

	const interval = 200 * time.Millisecond

	for {
		select {
		case <-ctx.Done():
			return "", nil, ctx.Err()
		case <-time.After(interval):
			resp, err := q.client.ReceiveMessageWithContext(ctx, input)
			if err != nil {
				return "", nil, err
			}

			if len(resp.Messages) > 0 {
				sqsMsg := resp.Messages[0]
				var msg message
				if sqsMsg.Body != nil {
					_ = jsoniter.UnmarshalFromString(*sqsMsg.Body, &msg)
				}

				callback := &sqsCallback{
					queue: q,
					msg:   sqsMsg,
				}

				if msg.IsExpired() {
					if err := callback.Finish(ctx); err != nil {
						return "", nil, err
					}

					break
				}

				// 还没到时间
				if d := time.Until(msg.VisibleAt); d >= time.Second {
					if err := callback.Delay(ctx, d); err != nil {
						return "", nil, err
					}

					break
				}

				return msg.Content, callback, nil
			}
		}
	}
}
