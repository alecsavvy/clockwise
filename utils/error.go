package utils

import (
	"errors"
	"fmt"
)

// utility to make joining errors easier
func AppError(msg string, err error) error {
	return errors.Join(errors.New(msg), err)
}

// utility to make joining errors easier but with formatting
// for additional error data
func AppErrorf(format string, err error, v ...any) error {
	msg := fmt.Sprintf(format, v...)
	return AppError(msg, err)
}
