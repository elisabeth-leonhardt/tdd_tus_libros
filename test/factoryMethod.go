package test

import (
	"TusLibros/source"
	"time"
)

type Factory struct {
}

func NewFactory() Factory {
	return Factory{}
}

// Factory Methods
func (f *Factory) ValidBook() source.Book {
	return source.Book{ISBN: "ValidBook", Price: 10.00}
}

func (f *Factory) InvalidBook() source.Book {
	return source.Book{ISBN: "InvalidBook"}
}

func (f *Factory) EmptyCart() source.Cart {
	bookValid := source.Book{ISBN: "ValidBook"}
	booksCatalogue := []source.Book{bookValid}
	return source.NewCart(booksCatalogue)
}

func (f *Factory) FullCartWithThreeBooks() source.Cart {
	bookValid := source.Book{ISBN: "ValidBook", Price: 10.00}
	booksCatalogue := []source.Book{bookValid}
	cart := source.NewCart(booksCatalogue)
	cart.AddObjectToCart(bookValid, 3)
	return cart
}

func (f *Factory) FullCartWithOneBook() source.Cart {
	bookValid := source.Book{ISBN: "ValidBook", Price: 10.00}
	booksCatalogue := []source.Book{bookValid}
	cart := source.NewCart(booksCatalogue)
	cart.AddObjectToCart(bookValid, 1)
	return cart
}

func (f *Factory) ValidCard() source.Card {
	numberCard := "12312312312"
	cardName := "Juan Perez"
	validFutureDate := time.Now().Add(time.Hour * 24)
	validCard := source.Card{
		Date:       &validFutureDate,
		NumberCard: &numberCard,
		Name:       &cardName,
	}
	return validCard
}

func (f *Factory) InvalidEmptyCard() source.Card {
	invalidCard := source.Card{
		Date:       nil,
		NumberCard: nil,
		Name:       nil,
	}
	return invalidCard
}

func (f *Factory) ExpiredCard() source.Card {
	numberCard := "12312312312"
	cardName := "Juan Perez"
	expiredTime := time.Date(2021, 01, 23, 0, 0, 0, 0, time.Local)
	invalidCard := source.Card{
		Date:       &expiredTime,
		NumberCard: &numberCard,
		Name:       &cardName,
	}
	return invalidCard
}

func (f *Factory) Cashier() source.Cashier {
	return source.Cashier{}
}

func (f *Factory) NewSuccessMerchantProcessor() source.MerchantProcessorStub {
	return source.MerchantProcessorStub{
		DebitBehavior: func(anAmount float64, aCreditCard source.Card) error {
			return nil
		},
	}
}

func (f *Factory) NewFailMerchantProcessor() source.MerchantProcessorStub {
	return source.MerchantProcessorStub{
		DebitBehavior: func(anAmount float64, aCreditCard source.Card) error {
			return source.Error06
		},
	}
}
