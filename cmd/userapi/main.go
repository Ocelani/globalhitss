package main

import (
	"globalhitss/internal/userapi"
	"globalhitss/pkg/infra/database"
	"os"
	"time"
)

const (
	DefaultPort = "3000" // Default port of this API

	DatabaseHost          = "localhost"
	DatabasePort          = 5432
	DatabaseUser          = "postgres"
	DatabasePassword      = "password123"
	DatabaseName          = "globalhitss"
	DatabaseEnableSSL     = false
	DatabaseTimezone      = "America/Sao_Paulo"
	DatabaseTimeout       = time.Second * 5
	DatabaseRetryInterval = time.Second
)

func getPort() (port string) {
	if port = os.Getenv("PORT"); port == "" {
		return DefaultPort
	}
	return port
}

func getDatabaseConnector() *database.Postgres {
	return &database.Postgres{
		Host:          DatabaseHost,
		Port:          DatabasePort,
		User:          DatabaseUser,
		Password:      DatabasePassword,
		NameDB:        DatabaseName,
		SSL:           DatabaseEnableSSL,
		Timezone:      DatabaseTimezone,
		Timeout:       DatabaseTimeout,
		RetryInterval: DatabaseTimeout,
	}
}

func main() {
	port := getPort()
	db := getDatabaseConnector()

	if err := db.Open(); err != nil {
		panic(err)
	}

	rep := userapi.NewUserRepository(db)
	svc := userapi.NewUserService(rep)
	hnd := userapi.NewUserHandler(svc)
	api := userapi.NewUserAPI(hnd)

	api.Setup()

	if err := api.Listen(port); err != nil {
		panic(err)
	}
}
