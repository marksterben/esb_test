package domain

import "errors"

var (
	ErrInternalServerError = errors.New("Internal Server Error")
	ErrBadParamInput       = errors.New("Given Params Is Invalid")
)
