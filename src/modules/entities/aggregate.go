package entities

import "github.com/google/uuid"

type UserAggregate struct {
	UserId uuid.UUID
}

func New(id uuid.UUID) *UserAggregate {
	return &UserAggregate{
		UserId: id,
	}
}
