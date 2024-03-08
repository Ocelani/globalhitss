package userapi

import "fmt"

const (
	BindBodyError    = "bind request body error: %w"
	UserHandlerError = "user handler error: %w"
)

func bindBodyError(err error) Response {
	return Response{
		Error: fmt.Errorf(BindBodyError, err).Error(),
	}
}

func userHandlerError(err error) Response {
	return Response{
		Error: fmt.Errorf(UserHandlerError, err).Error(),
	}
}
