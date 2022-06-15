package source

import (
	"errors"
	"time"
)

var Error03 = errors.New("cannot checkout a empty cart")
var Error04 = errors.New("cannot checkout a invalid card")
var Error05 = errors.New("cannot checkout a expired card")

type Cashier struct {
}

type Card struct {
	Date       *time.Time
	NumberCard *string
	Name       *string
}

func (c *Cashier) Checkout(aCart Cart, aCard Card, merchant MerchantProcessorStub) (total float64, err error) {
	if len(aCart.Books) <= 0 {
		return 0, Error03
	}

	if aCard.Name == nil || aCard.NumberCard == nil || aCard.Date == nil {
		return 0, Error04
	}

	if aCard.Date.Before(time.Now()) {
		return 0, Error05
	}

	for _, b := range aCart.Books {
		total = total + b.Price
	}

	err = merchant.ProcessPayment(aCard, total)

	return total, err
}
