package entities

type UserAggregate struct {
	UserId string
}

func New(id string) *UserAggregate {
	return &UserAggregate{
		UserId: id,
	}
}
