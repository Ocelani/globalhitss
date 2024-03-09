package main

import (
	"context"
	"globalhitss/pkg/infra/database"
	"globalhitss/pkg/infra/queue"
	"os"
	"time"
)

const (
	DefaultPort = "3000"

	DatabaseHost          = "localhost"
	DatabasePort          = 5432
	DatabaseUser          = "postgres"
	DatabasePassword      = "password123"
	DatabaseName          = "globalhitss"
	DatabaseEnableSSL     = false
	DatabaseTimezone      = "America/Sao_Paulo"
	DatabaseTimeout       = time.Second * 5
	DatabaseRetryInterval = time.Second

	QueueUser     = "guest"
	QueuePassword = "guest"
	QueueHost     = "localhost"
	QueuePort     = 5672

	QueueName             = "user"
	QueueDurable          = false
	QueueDeleteWhenUnused = false
	QueueAutoDelete       = false
	QueueExclusive        = false
	QueueNoWait           = false

	QueueExchangeKind     = "direct"
	QueueExchangeInternal = false
)

// getPort returns the port to be used by the API.
func getPort() (port string) {
	if port = os.Getenv("PORT"); port == "" {
		return DefaultPort
	}
	return port
}

// getDatabaseConfig returns the configuration for the database.
func getDatabaseConfig() *database.Postgres {
	return &database.Postgres{
		Host:          DatabaseHost,
		Port:          DatabasePort,
		User:          DatabaseUser,
		Password:      DatabasePassword,
		NameDB:        DatabaseName,
		SSL:           DatabaseEnableSSL,
		Timezone:      DatabaseTimezone,
		Timeout:       DatabaseTimeout,
		RetryInterval: DatabaseRetryInterval,
	}
}

// getQueueConfig returns the configuration for the queue.
func getQueueConfig() *queue.Config {
	return &queue.Config{
		Queue: &queue.ConfigQueue{
			Name:             QueueName,
			Durable:          QueueDurable,
			DeleteWhenUnused: QueueDeleteWhenUnused,
			AutoDelete:       QueueAutoDelete,
			Exclusive:        QueueExclusive,
			NoWait:           QueueNoWait,
		},
		Exchange: &queue.ConfigExchange{
			Name:       QueueName,
			Kind:       QueueExchangeKind,
			Durable:    QueueDurable,
			AutoDelete: QueueAutoDelete,
			Internal:   QueueExchangeInternal,
			NoWait:     QueueNoWait,
		},
		Dial: &queue.ConfigDial{
			User:     QueueUser,
			Password: QueuePassword,
			Host:     QueueHost,
			Port:     QueuePort,
		},
	}
}

// newQueueService returns a new instance of the queue service.
func newQueueService(repository *UserRepository) (*UserQueueService, error) {
	cfg := getQueueConfig()

	prd, err := queue.NewProducer(cfg)
	if err != nil {
		return nil, err
	}

	cns, err := queue.NewConsumer(cfg)
	if err != nil {
		return nil, err
	}

	return NewUserQueueService(
		NewUserQueueConsumer(cns, repository),
		NewUserQueueProducer(prd),
	), nil
}

// main is the entrypoint of the application.
func main() {
	port := getPort()
	db := getDatabaseConfig()

	if err := db.Open(); err != nil {
		panic(err)
	}

	repository := NewUserRepository(db)

	queueService, err := newQueueService(repository)
	if err != nil {
		panic(err)
	}

	go queueService.ConsumeCreate(context.Background())

	service := NewUserService(repository)
	handler := NewUserHandler(service, queueService)

	api := NewUserAPI(handler)
	api.Setup()

	if err := api.Listen(port); err != nil {
		panic(err)
	}
}
