package queue

import (
	"context"

	amqp "github.com/rabbitmq/amqp091-go"
)

// Producer for AMQP.
type Producer struct {
	Channel *amqp.Channel
	Conn    *amqp.Connection
	queue   *amqp.Queue
	config  *Config
}

// NewProducer for AMQP.
func NewProducer(cfg *Config) (*Producer, error) {
	conn, err := amqp.Dial(cfg.Dial.String())
	if err != nil {
		return nil, err
	}
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}
	if err = exchangeDeclare(ch, cfg.Exchange); err != nil {
		return nil, err
	}
	queue, err := queueDeclare(ch, cfg.Queue)
	if err != nil {
		return nil, err
	}

	return &Producer{
		Conn:    conn,
		Channel: ch,
		queue:   &queue,
		config:  cfg,
	}, nil
}

// Push (Publish) a specified message to the AMQP exchange.
func (p *Producer) Publish(ctx context.Context, body []byte) error {
	return p.Channel.PublishWithContext(
		ctx,
		p.config.Exchange.Name,
		p.config.Exchange.Key,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		},
	)
}
