package vo

import (
	"fmt"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type Password struct {
	Value []byte
}

type PasswordError struct {
	ErrorMsg string
}

func (e PasswordError) Error() string {
	return fmt.Sprintln(e.ErrorMsg)
}

func CreatePassword(password string) (Password, error) {
	if len(strings.TrimSpace(password)) == 0 {
		return Password{}, PasswordError{ErrorMsg: "Password has to be provided"}
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		fmt.Println(err)
		return Password{}, PasswordError{ErrorMsg: "Password encyption failed"}
	}
	return newPassword(hash), nil
}

func VerifyPassword(password string, hashedPassword []byte) error {
	err := bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	if err != nil {
		return PasswordError{ErrorMsg: "You have provided an invalid password"}
	}
	return nil
}

func newPassword(password []byte) Password {
	return Password{
		Value: password,
	}
}
