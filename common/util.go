package common

import (
	"context"
	"errors"
)

// IsCtxCanceled checks if an error is caused by context.Canceled
func IsCtxCanceled(err error) bool {
	if err == nil {
		return false
	}

	return errors.Is(err, context.Canceled)
}

// NonCtxCanceledErr returns nil if the error is caused by context.Canceled
func NonCtxCanceledErr(err error) error {
	if IsCtxCanceled(err) {
		return nil
	}

	return err
}
