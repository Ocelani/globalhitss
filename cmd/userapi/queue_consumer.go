package main

import (
	"context"
	"encoding/json"
	"fmt"
	"globalhitss/pkg/infra/queue"

	"github.com/gofiber/fiber/v3/log"
	"github.com/rabbitmq/amqp091-go"
)

// UserQueueConsumer consumes messages from the user-create queue.
type UserQueueConsumer struct {
	Consumer   *queue.Consumer
	Repository *UserRepository
}

// NewUserQueueConsumer returns a new UserQueueConsumer.
func NewUserQueueConsumer(consumer *queue.Consumer, repository *UserRepository) *UserQueueConsumer {
	return &UserQueueConsumer{
		Consumer:   consumer,
		Repository: repository,
	}
}

// Consume consumes messages from the user-create queue.
func (u *UserQueueConsumer) ConsumeCreate(ctx context.Context) {
	msgCH, err := u.Consumer.Consume(ctx, false, false)
	if err != nil {
		log.Fatal(err)
	}
	for {
		if err := u.handleMessage(ctx, <-msgCH); err != nil {
			log.Error(err)
			continue
		}
	}
}

// consumeMessages consumes messages from the user-create queue.
func (u *UserQueueConsumer) handleMessage(ctx context.Context, msg amqp091.Delivery) error {
	var user User

	if err := json.Unmarshal(msg.Body, &user); err != nil {
		return fmt.Errorf("failed to unmarshal message: %w", err)
	}

	if err := u.Repository.Create(ctx, toUserEntity(user)); err != nil {
		return fmt.Errorf("failed to create user on repository: %w", err)
	}

	if err := msg.Ack(false); err != nil {
		return fmt.Errorf("failed to acknowledge message: %w", err)
	}

	return nil
}
