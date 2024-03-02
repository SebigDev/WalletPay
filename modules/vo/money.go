package vo

type Money struct {
	Amount   Amount
	Currency Currency
}

func NewMoney(amount Amount, currency Currency) *Money {
	return &Money{
		Amount:   amount,
		Currency: currency,
	}
}
