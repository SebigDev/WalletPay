package utils

import "fmt"

type AppError struct {
	ErrorMsg string
}

func (e AppError) Error() string {
	return fmt.Sprintf(e.ErrorMsg)
}
