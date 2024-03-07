package main

import (
	"globalhitss/internal/api"
	"os"
	"time"
)

const (
	defaultPort = "3000"          // Default port of this API
	dbTimeout   = time.Second * 5 // Timeout to connect to database
)

func getPort() (port string) {
	port = os.Getenv("PORT")
	if port == "" {
		return defaultPort
	}
	return port
}

func main() {
	port := getPort()
	a := api.NewAPI()

	defer a.Close()
	a.Setup()

	if err := a.Listen(port); err != nil {
		panic(err)
	}
}
