package sqs

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type sqsCallback struct {
	*queue
	msg *sqs.Message
}

func (s *sqsCallback) Finish(ctx context.Context) error {
	_, err := s.client.DeleteMessage(&sqs.DeleteMessageInput{
		QueueUrl:      s.queueURL,
		ReceiptHandle: s.msg.ReceiptHandle,
	})

	return err
}

func (s *sqsCallback) Delay(ctx context.Context, interval time.Duration) error {
	if interval > 12*time.Hour {
		// sqs 最大限制
		interval = 12 * time.Hour
	}

	seconds := interval.Milliseconds() / 1000
	if seconds == 0 {
		seconds = 1
	}

	_, err := s.client.ChangeMessageVisibility(&sqs.ChangeMessageVisibilityInput{
		QueueUrl:          s.queueURL,
		ReceiptHandle:     s.msg.ReceiptHandle,
		VisibilityTimeout: aws.Int64(seconds),
	})

	return err
}
