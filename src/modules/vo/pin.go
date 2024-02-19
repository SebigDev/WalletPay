package vo

import (
	"fmt"

	"github.com/google/uuid"
)

type PinError struct {
	ErrorMsg string
}

func (e PinError) Error() string {
	return fmt.Sprintf(e.ErrorMsg)
}

type Value string

func (p *Value) String() string {
	return string(*p)
}

func NewValue(val string) (Value, error) {
	if len(val) == 0 {
		return "", PinError{ErrorMsg: "Pin value must be provided"}
	}
	if len(val) < 4 || len(val) > 4 {
		return "", PinError{ErrorMsg: "Pin value must have a length of 4"}
	}
	return Value(val), nil
}

type Pin struct {
	ValueHash    []byte
	RecoverValue string
}

func NewPin(pin Value) *Pin {
	return &Pin{
		ValueHash:    []byte(pin),
		RecoverValue: uuid.NewString(),
	}
}

func Verify(pin string, p Pin) error {
	hashedPin := []byte(pin)
	for i, b := range p.ValueHash {
		if b != hashedPin[i] {
			return PinError{ErrorMsg: "You have provided an invalid PIN"}
		}
	}
	return nil
}

func VerifyRecover(recoveryPin string, p Pin) error {
	if recoveryPin != p.RecoverValue {
		return PinError{ErrorMsg: "You have provided an invalid recovery PIN"}
	}
	return nil
}
