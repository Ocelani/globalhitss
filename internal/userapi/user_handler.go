package userapi

import (
	"net/http"

	"github.com/gofiber/fiber/v3"
)

const (
	getUserRoute    = "/user/:id" // Route to get a user by ID.
	postUserRoute   = "/user"     // Route to create a new user.
	putUserRoute    = "/user/:id" // Route to update a user by ID.
	deleteUserRoute = "/user/:id" // Route to delete a user by ID.
)

type UserHandler struct {
	Service *UserService
}

func NewUserHandler(service *UserService) *UserHandler {
	return &UserHandler{
		Service: service,
	}
}

// routes set up the API routes for the user.
func (h *UserHandler) routes(api *UserAPI) {
	api.Server.Get(getUserRoute, h.Get)
	api.Server.Post(postUserRoute, h.Post)
	api.Server.Put(putUserRoute, h.Put)
	api.Server.Delete(deleteUserRoute, h.Delete)
}

// Get returns a user by ID.
func (h *UserHandler) Get(c fiber.Ctx) error {
	return c.Status(http.StatusOK).SendString("get user")
}

// Post creates a new user.
func (h *UserHandler) Post(c fiber.Ctx) error {
	return c.Status(http.StatusOK).SendString("post user")
}

// Put updates a user by ID.
func (h *UserHandler) Put(c fiber.Ctx) error {
	return c.Status(http.StatusOK).SendString("put user")
}

// Delete deletes a user by ID.
func (h *UserHandler) Delete(c fiber.Ctx) error {
	return c.Status(http.StatusOK).SendString("delete user")
}
