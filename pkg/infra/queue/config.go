package queue

import (
	"fmt"
	"net"
	"strconv"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Config struct {
	Exchange *ConfigExchange
	Queue    *ConfigQueue
	Dial     *ConfigDial
}

// ConfigExchange for AMQP.
type ConfigExchange struct {
	Name       string
	Key        string
	Kind       string
	Durable    bool
	AutoDelete bool
	Internal   bool
	NoWait     bool
	Args       amqp.Table
}

// ConfigQueue for AMQP.
type ConfigQueue struct {
	Name             string
	Durable          bool
	DeleteWhenUnused bool
	AutoDelete       bool
	Exclusive        bool
	NoWait           bool
	Args             amqp.Table
}

// ConfigDial for AMQP.
type ConfigDial struct {
	User     string
	Password string
	Host     string
	Port     int
}

func (c *ConfigDial) String() string {
	hostport := net.JoinHostPort(c.Host, strconv.Itoa(c.Port))
	return fmt.Sprintf("amqp://%s:%s@%s", c.User, c.Password, hostport)
}

func exchangeDeclare(ch *amqp.Channel, cfg *ConfigExchange) error {
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
