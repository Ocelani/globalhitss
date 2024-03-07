package api

import (
	"fmt"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
)

// API is the implementation of the repository interface.
type API struct {
	Server *fiber.App
	user   *User
}

// NewAPI returns a new instance of the default repository.
func NewAPI() *API {
	return &API{
		Server: fiber.New(),
		user:   NewUser(),
	}
}

// Setup API service.
func (a *API) Setup() {
	a.Server.Use(logger.New(logger.ConfigDefault))

	a.Server.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin, Content-Type, Accept, X-Auth",
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowOrigins:     "*",
		AllowCredentials: true,
	}))

	a.Server.Get("/", func(http fiber.Ctx) error {
		http.Accepts("application/json", "text/html", "html", "text", "json")
		return http.SendString("Backend Challenge - Global Hitss")
	})

	a.user.routes(a)
}

// Listen and serve the API service.
func (a *API) Listen(port string) error {
	return a.Server.Listen(
		fmt.Sprintf(":%s", port),
	)
}

// Close API service.
func (a *API) Close() {
	if err := a.Server.Shutdown(); err != nil {
		panic(err)
	}
}
