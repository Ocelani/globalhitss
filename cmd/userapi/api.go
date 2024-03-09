package main

import (
	"fmt"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
)

const (
	getUserRoute       = "/user/:id"   // Route to get a user by ID.
	postUserRoute      = "/user"       // Route to create a new user.
	putUserRoute       = "/user/:id"   // Route to update a user by ID.
	deleteUserRoute    = "/user/:id"   // Route to delete a user by ID.
	postUserQueueRoute = "/user/queue" // Route to create a new user by queue.
)

// API is the implementation of the repository interface.
type API struct {
	Server *fiber.App
	user   *UserHandler
}

// NewUserAPI returns a new instance of the default repository.
func NewUserAPI(userHandler *UserHandler) *API {
	return &API{
		Server: fiber.New(),
		user:   userHandler,
	}
}

// Setup API service.
func (a *API) Setup() {
	a.Server.Use(logger.New(logger.ConfigDefault))

	a.Server.Use(cors.New(cors.Config{
		AllowHeaders: "Origin, Content-Type, Accept, X-Auth",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH",
	}))

	a.Server.Get("/", func(http fiber.Ctx) error {
		http.Accepts("application/json", "text/html", "html", "text", "json")
		return http.SendString("Backend Challenge - Global Hitss\nAPI User Service\n")
	})

	a.user.routes(a)
}

// Listen and serve the API service.
func (a *API) Listen(port string) error {
	defer a.close()
	return a.Server.Listen(fmt.Sprintf(":%s", port))
}

// close API service.
func (a *API) close() {
	if err := a.Server.Shutdown(); err != nil {
		panic(err)
	}
}
