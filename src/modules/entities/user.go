package entities

import (
	"CrashCourse/GoApp/src/modules/vo"
	"time"
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
}
