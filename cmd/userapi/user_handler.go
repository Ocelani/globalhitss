package main

import (
	"encoding/json"
	"globalhitss/pkg/user"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v3"
)

// UserHandler is the handler for the user.
type UserHandler struct {
	Service *UserService
	Queue   *UserQueueService
}

// NewUserHandler returns a new instance of the default repository.
func NewUserHandler(service *UserService, queue *UserQueueService) *UserHandler {
	return &UserHandler{
		Service: service,
		Queue:   queue,
	}
}

// routes set up the API routes for the user.
func (h *UserHandler) routes(api *API) {
	api.Server.Get(getUserRoute, h.Get)
	api.Server.Post(postUserRoute, h.Post)
	api.Server.Put(putUserRoute, h.Put)
	api.Server.Delete(deleteUserRoute, h.Delete)
	api.Server.Post(postUserQueueRoute, h.PostQueue)
}

// Get returns a user by ID.
func (h *UserHandler) Get(ctx fiber.Ctx) error {
	id, err := strconv.ParseUint(ctx.Params("id"), 10, 64)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(userHandlerError(err))
	}

	data, err := h.Service.ReadOne(ctx.Context(), user.ID(id))
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(userHandlerError(err))
	}

	return ctx.Status(http.StatusOK).JSON(toResponse(data))
}

// Post creates a new user.
func (h *UserHandler) Post(ctx fiber.Ctx) error {
	var dto User
	if err := json.Unmarshal(ctx.Body(), &dto); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(userHandlerError(err))
	}

	ent := toUserEntity(dto)
	if err := h.Service.Create(ctx.Context(), ent); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(userHandlerError(err))
	}

	return ctx.Status(http.StatusOK).JSON(toResponse(ent))
}

// Put updates a user by ID.
func (h *UserHandler) Put(ctx fiber.Ctx) error {
	id, err := strconv.ParseUint(ctx.Params("id"), 10, 64)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(userHandlerError(err))
	}

	var dto User
	if err := json.Unmarshal(ctx.Body(), &dto); err != nil {
		_ = ctx.Status(http.StatusBadRequest).JSON(bindBodyError(err))
		return err
	}

	ent := toUserEntity(dto)
	if err := h.Service.Update(ctx.Context(), user.ID(id), ent); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(userHandlerError(err))
	}

	return ctx.Status(http.StatusOK).JSON(toResponse(ent))
}

// Delete deletes a user by ID.
func (h *UserHandler) Delete(ctx fiber.Ctx) error {
	id, err := strconv.ParseUint(ctx.Params("id"), 10, 64)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(userHandlerError(err))
	}

	if err = h.Service.Delete(ctx.Context(), user.ID(id)); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(userHandlerError(err))
	}

	return ctx.Status(http.StatusOK).JSON(User{
		ID: uint(id),
	})
}

func (h *UserHandler) PostQueue(ctx fiber.Ctx) error {
	var dto User
	if err := json.Unmarshal(ctx.Body(), &dto); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(userHandlerError(err))
	}

	if err := h.Queue.PublishCreate(ctx.Context(), &dto); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(userHandlerError(err))
	}

	return ctx.Status(http.StatusOK).JSON(dto)
}
