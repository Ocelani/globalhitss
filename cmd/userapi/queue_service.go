package main

import (
	"context"
)

type UserQueueService struct {
	Producer *UserQueueProducer
	Consumer *UserQueueConsumer
}

func NewUserQueueService(
	consumer *UserQueueConsumer,
	producer *UserQueueProducer,
) *UserQueueService {
	return &UserQueueService{
		Consumer: consumer,
		Producer: producer,
	}
}

func (u *UserQueueService) ConsumeCreate(ctx context.Context) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	u.Consumer.ConsumeCreate(ctx)
}

func (u *UserQueueService) PublishCreate(ctx context.Context, data *User) error {
	return u.Producer.PublishCreate(ctx, data)
}
