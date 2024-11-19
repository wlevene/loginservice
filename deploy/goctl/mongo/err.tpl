package model

import "errors"

var (
	ErrNotFound        = errors.New("Not Found")
	ErrInvalidObjectId = errors.New("invalid objectId")
	ErrNilObj          = errors.New("nil object")
)
