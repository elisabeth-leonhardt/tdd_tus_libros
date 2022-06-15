package test

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type CheckoutTestSuite struct {
	suite.Suite
}

func (suite *CheckoutTestSuite) TestCannotCheckoutAndEmptyCart() {
	aCart := EmptyCart()
	cantCart := len(aCart.Books)
	suite.Equal(0, cantCart)
}

func TestRunCheckoutSuite(t *testing.T) {
	suite.Run(t, new(CheckoutTestSuite))
}
