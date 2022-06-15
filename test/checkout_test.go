package test

import (
	"TusLibros/source"
	"github.com/stretchr/testify/suite"
	"testing"
)

type CheckoutTestSuite struct {
	suite.Suite
	factory Factory
}

func (suite *CheckoutTestSuite) SetupTest() {
	suite.factory = NewFactory()
}

func (suite *CheckoutTestSuite) TestCannotCheckoutAndEmptyCart() {
	aCart := suite.factory.EmptyCart()
	aCashier := suite.factory.Cashier()
	aValidCard := suite.factory.ValidCard()
	_, err := aCashier.Checkout(aCart, aValidCard)
	suite.Equal(err, source.Error03)
}

func (suite *CheckoutTestSuite) TestValidCard() {
	aCart := suite.factory.FullCart()
	aCashier := suite.factory.Cashier()
	aValidCard := suite.factory.ValidCard()
	total, _ := aCashier.Checkout(aCart, aValidCard)
	suite.Equal(total, 1.00)
}

func (suite *CheckoutTestSuite) TestInvalidCard() {
	aCart := suite.factory.FullCart()
	aCashier := suite.factory.Cashier()
	aInvalidCard := suite.factory.InvalidEmptyCard()
	_, err := aCashier.Checkout(aCart, aInvalidCard)
	suite.Equal(err, source.Error04)
}

func (suite *CheckoutTestSuite) TestInvalidCardDate() {
	aCart := suite.factory.FullCart()
	aCashier := suite.factory.Cashier()
	aExpiredCard := suite.factory.ExpiredCard()
	_, err := aCashier.Checkout(aCart, aExpiredCard)
	suite.Equal(err, source.Error05)
}

func TestRunCheckoutSuite(t *testing.T) {
	suite.Run(t, new(CheckoutTestSuite))
}
