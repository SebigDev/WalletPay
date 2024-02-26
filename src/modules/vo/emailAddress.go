package vo

import (
	"CrashCourse/GoApp/internal/utils"
	"fmt"
	"net/mail"
)

type EmailAddress struct {
	Value string
}

type EmailAddressError struct {
	ErrorMsg string
}

func (e EmailAddressError) Error() string {
	return fmt.Sprintln(e.ErrorMsg)
}

func CreateEmailAddress(emailAddress string) (*EmailAddress, error) {
	if utils.Length(emailAddress) == 0 {
		return &EmailAddress{}, EmailAddressError{ErrorMsg: "Email address has to be provided"}
	}
	isValid := validateEmail(emailAddress)
	if !isValid {

		return &EmailAddress{}, EmailAddressError{ErrorMsg: "You have provided an invalid email address"}
	}
	return newEmailAddress(emailAddress), nil
}

func newEmailAddress(emailAddress string) *EmailAddress {
	return &EmailAddress{
		Value: emailAddress,
	}
}

func validateEmail(emailAddress string) bool {
	_, err := mail.ParseAddress(emailAddress)
	return err == nil
}
