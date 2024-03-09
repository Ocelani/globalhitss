package main

import (
	"context"
	"time"
)

// UserQueueService provides the business logic for the user entity.
type UserQueueService struct {
	Producer *UserQueueProducer
	Consumer *UserQueueConsumer
}

// NewUserQueueService returns a new UserQueueService.
func NewUserQueueService(
	consumer *UserQueueConsumer,
	producer *UserQueueProducer,
) *UserQueueService {
	return &UserQueueService{
		Consumer: consumer,
		Producer: producer,
	}
}

// ConsumeCreate consumes messages from the user-create queue.
func (u *UserQueueService) ConsumeCreate(ctx context.Context) {
	for {
		u.Consumer.ConsumeCreate(ctx)
		time.Sleep(1 * time.Second)
	}
}

// PublishCreate publishes a message to the user-create queue to create a new user.
func (u *UserQueueService) PublishCreate(ctx context.Context, data *User) error {
	return u.Producer.PublishCreate(ctx, data)
}
