package types

import "github.com/wlevene/loginservice/internal/errorx"

func ErrPermissionDenied() error {
	return errorx.NewCodeError(417, "permission denied")
}

func ErrInvalidParameter() error {
	return errorx.NewCodeError(470, "invalid parameter")
}
