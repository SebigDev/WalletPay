package vo

import (
	"fmt"
	"net/mail"

	"github.com/sebigdev/walletpay/internal/utils"
)

type EmailAddress struct {
	Value string
}

type EmailAddressError struct {
	ErrorMsg string
}

func (e EmailAddressError) Error() string {
	return fmt.Sprintf(e.ErrorMsg)
}

func CreateEmailAddress(emailAddress string) (*EmailAddress, error) {
	if utils.Length(emailAddress) == 0 {
		return &EmailAddress{}, EmailAddressError{ErrorMsg: "Email address has to be provided"}
	}
	isValid := validateEmail(emailAddress)
	if !isValid {

		return &EmailAddress{}, EmailAddressError{ErrorMsg: "You have provided an invalid email address"}
	}
	return &EmailAddress{
		Value: emailAddress,
	}, nil
}

func validateEmail(emailAddress string) bool {
	_, err := mail.ParseAddress(emailAddress)
	return err == nil
}
