package api

import (
	"github.com/gofiber/fiber/v3"
)

const (
	getUserRoute    = "/user/:id" // Route to get a user by ID.
	postUserRoute   = "/user"     // Route to create a new user.
	putUserRoute    = "/user/:id" // Route to update a user by ID.
	deleteUserRoute = "/user/:id" // Route to delete a user by ID.
)

type User struct {
}

func NewUser() *User {
	return &User{}
}

// setupUserRoutes sets up the routes for the user.
func (u *User) routes(api *API) {
	api.Server.Get(getUserRoute, u.get)
	api.Server.Post(postUserRoute, u.post)
	api.Server.Put(putUserRoute, u.put)
	api.Server.Delete(deleteUserRoute, u.delete)
}

// get returns a user by ID.
func (u *User) get(http fiber.Ctx) error {
	return http.JSON(fiber.Map{})
}

// post creates a new user.
func (u *User) post(http fiber.Ctx) error {
	return http.JSON(fiber.Map{})
}

// put updates a user by ID.
func (u *User) put(http fiber.Ctx) error {
	return http.JSON(fiber.Map{})
}

// delete deletes a user by ID.
func (u *User) delete(http fiber.Ctx) error {
	return http.JSON(fiber.Map{})
}
