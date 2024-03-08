package main

import (
	"context"
	"encoding/json"
	"globalhitss/pkg/infra/queue"
)

// UserQueueProducer publishes messages to the user-create queue.
type UserQueueProducer struct {
	Producer *queue.Producer
}

// NewUserQueueProducer returns a new UserQueueProducer.
func NewUserQueueProducer(producer *queue.Producer) *UserQueueProducer {
	return &UserQueueProducer{
		Producer: producer,
	}
}

// PublishCreate publishes a message to the user-create queue to create a new user.
func (u *UserQueueProducer) PublishCreate(ctx context.Context, data *User) error {
	buf, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return u.Producer.Publish(ctx, "user.create", buf)
}
