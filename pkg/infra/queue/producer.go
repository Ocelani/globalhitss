package queue

import (
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
func (p *Producer) Publish(body []byte, key string) error {
	return p.Channel.Publish(
		p.config.Exchange.Name,
		key,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		},
	)
}

func declareExchange(ch *amqp.Channel, cfg *ConfigQueue) error {
	return ch.ExchangeDeclare(
		cfg.Name,
		cfg.Kind,
		cfg.Durable,
		cfg.AutoDelete,
		cfg.Internal,
		cfg.NoWait,
		cfg.Args,
	)
}

func queueDeclare(ch *amqp.Channel, cfg *ConfigQueue) (amqp.Queue, error) {
	return ch.QueueDeclare(
		cfg.Name,
		cfg.Durable,
		cfg.AutoDelete,
		cfg.Exclusive,
		cfg.NoWait,
		cfg.Args,
	)
}
