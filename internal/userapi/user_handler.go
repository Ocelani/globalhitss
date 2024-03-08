package userapi

import (
	"encoding/json"
	"globalhitss/pkg/user"
	"net/http"
	"strconv"

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
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(userHandlerError(err))
	}

	data, err := h.Service.ReadOne(c.Context(), user.ID(id))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(userHandlerError(err))
	}

	return c.Status(http.StatusOK).JSON(toResponse(data))
}

// Post creates a new user.
func (h *UserHandler) Post(c fiber.Ctx) error {
	var dto User
	if err := json.Unmarshal(c.Body(), &dto); err != nil {
		return c.Status(http.StatusBadRequest).JSON(userHandlerError(err))
	}

	ent := toUserEntity(dto)
	if err := h.Service.Create(c.Context(), ent); err != nil {
		return c.Status(http.StatusBadRequest).JSON(userHandlerError(err))
	}

	return c.Status(http.StatusOK).JSON(toResponse(ent))
}

// Put updates a user by ID.
func (h *UserHandler) Put(c fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(userHandlerError(err))
	}

	var dto User
	if err := json.Unmarshal(c.Body(), &dto); err != nil {
		_ = c.Status(http.StatusBadRequest).JSON(bindBodyError(err))
		return err
	}

	ent := toUserEntity(dto)
	if err := h.Service.Update(c.Context(), user.ID(id), ent); err != nil {
		return c.Status(http.StatusBadRequest).JSON(userHandlerError(err))
	}

	return c.Status(http.StatusOK).JSON(toResponse(ent))
}

// Delete deletes a user by ID.
func (h *UserHandler) Delete(c fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(userHandlerError(err))
	}

	if err = h.Service.Delete(c.Context(), user.ID(id)); err != nil {
		return c.Status(http.StatusBadRequest).JSON(userHandlerError(err))
	}

	return c.Status(http.StatusOK).JSON(User{
		ID: uint(id),
	})
}
