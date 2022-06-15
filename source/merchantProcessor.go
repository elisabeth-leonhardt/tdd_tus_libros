package source

import "errors"

type MerchantProcessorStub struct {
	DebitBehavior func(anAmount float64, aCreditCard Card) error
}

var Error06 = errors.New("fail process merchant")

func NewMerchantProcessorStub(aDebitBehavior func(anAmount float64, aCreditCard Card) error) MerchantProcessorStub {
	return MerchantProcessorStub{aDebitBehavior}
}

func (m *MerchantProcessorStub) ProcessPayment(card Card, total float64) error {
	err := m.DebitBehavior(total, card)
	return err
}
