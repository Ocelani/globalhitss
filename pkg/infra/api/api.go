package api

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
)

// API is the implementation of the repository interface.
type API struct {
	Server *fiber.App
}

// NewAPI returns a new instance of the default repository.
func NewAPI() *API {
	return &API{
		Server: fiber.New(),
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
}

// Listen and serve the API service.
func (a *API) Listen(port string) error {
	defer a.close()
	return a.Server.Listen(":" + port)
}

// close API service.
func (a *API) close() {
	if err := a.Server.Shutdown(); err != nil {
		panic(err)
	}
}
