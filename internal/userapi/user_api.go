package userapi

import (
	"fmt"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
)

// UserAPI is the implementation of the repository interface.
type UserAPI struct {
	Server *fiber.App
	user   *UserHandler
}

// NewUserAPI returns a new instance of the default repository.
func NewUserAPI(userHandler *UserHandler) *UserAPI {
	return &UserAPI{
		Server: fiber.New(),
		user:   userHandler,
	}
}

// Setup API service.
func (a *UserAPI) Setup() {
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
func (a *UserAPI) Listen(port string) error {
	defer a.close()
	return a.Server.Listen(
		fmt.Sprintf(":%s", port),
	)
}

// close API service.
func (a *UserAPI) close() {
	if err := a.Server.Shutdown(); err != nil {
		panic(err)
	}
}
