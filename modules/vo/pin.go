package vo

import (
	"fmt"
	"strings"

	"github.com/sebigdev/walletpay/internal/utils"

	"github.com/google/uuid"
)

type PinError struct {
	ErrorMsg string
}

func (e PinError) Error() string {
	return fmt.Sprintf(e.ErrorMsg)
}

type PinValue string

func (p *PinValue) String() string {
	return string(*p)
}

func NewPinValue(val string) (PinValue, error) {
	if utils.Length(val) == 0 {
		return "", PinError{ErrorMsg: "Pin value must be provided"}
	}
	if utils.Length(val) < 4 || utils.Length(val) > 4 {
		return "", PinError{ErrorMsg: "Pin value must have a length of 4"}
	}
	if strings.Contains(val, "0000") {
		return "", PinError{ErrorMsg: "Pin value must have a length of 4"}
	}
	return PinValue(val), nil
}

type Pin struct {
	HashValue    []byte
	RecoverValue string
}

func NewPin(pin PinValue) *Pin {
	return &Pin{
		HashValue:    []byte(pin),
		RecoverValue: uuid.NewString(),
	}
}

func Verify(pin string, p Pin) error {
	savedPin := string(p.HashValue)
	cleanPin, err := NewPinValue(pin)
	if err != nil {
		return err
	}

	if savedPin == cleanPin.String() {
		return nil
	}
	return PinError{ErrorMsg: "You have provided an invalid PIN"}
}

func VerifyRecover(recoveryPin string, p Pin) error {
	if utils.Length(recoveryPin) == 0 {
		return PinError{ErrorMsg: "You have provided an invalid or empty recovery PIN"}
	}
	if recoveryPin != p.RecoverValue {
		return PinError{ErrorMsg: "You have provided an invalid recovery PIN"}
	}
	return nil
}
