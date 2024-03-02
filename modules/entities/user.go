package entities

import (
	"time"

	"github.com/sebigdev/walletpay/modules/vo"
)

type User struct {
	Aggregate    *UserAggregate
	FirstName    string
	LastName     string
	EmailAddress vo.EmailAddress
	Password     vo.Password
	CreateAt     time.Time
	IsActive     bool
	IsVerified   bool
	Pin          vo.Pin
}
