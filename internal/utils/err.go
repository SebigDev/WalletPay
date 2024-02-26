package utils

import "fmt"

type AppError struct {
	ErrorMsg string
}

func (e AppError) Error() string {
	return fmt.Sprintf(e.ErrorMsg)
}

func Check(err error) error {
	if err != nil {
		return err
	}
	return nil
}

func CheckWithMessage(msg string, err error) error {
	if err != nil {
		return fmt.Errorf("%v : %s", msg, err.Error())
	}
	return nil
}
