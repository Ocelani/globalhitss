package queue

import (
	"context"

	amqp "github.com/rabbitmq/amqp091-go"
)

// Consumer for AMQP.
type Consumer struct {
	Channel *amqp.Channel
	queue   *amqp.Queue
	Conn    *amqp.Connection
	config  *Config
}

// NewConsumer for AMQP.
func NewConsumer(cfg *Config) (*Consumer, error) {
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

	consumer := &Consumer{
		Channel: ch,
		Conn:    conn,
		queue:   &queue,
		config:  cfg,
	}

	err = consumer.queueBind(cfg.Exchange.Key, cfg.Exchange.Name)

	return consumer, err
}

func (c *Consumer) queueBind(key, exchange string) error {
	return c.Channel.QueueBind(
		c.queue.Name,
		key,
		exchange,
		c.config.Queue.NoWait,
		c.config.Queue.Args,
	)
}

func (c *Consumer) Consume(ctx context.Context, autoAck, noLocal bool) (<-chan amqp.Delivery, error) {
	return c.Channel.ConsumeWithContext(
		ctx,
		c.queue.Name,
		c.config.Exchange.Key,
		autoAck,
		c.config.Queue.Exclusive,
		noLocal,
		c.config.Queue.NoWait,
		c.config.Queue.Args,
	)
}
