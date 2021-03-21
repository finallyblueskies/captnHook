package rest

import "net/http"

const (
	// BadRequest is an http 400 response
	ErrBadRequest = http.StatusBadRequest
	// BodyBindingErr is an error message when we fail to bind the request body
	ErrBodyBinding = "could not bind body"
	// Internal Server err is a http 500 response
	ErrInternalServer = http.StatusInternalServerError
	// ErrStringConversion is an error message when we fail to convert a string
	ErrStringConversion = "could not convert string"
)

