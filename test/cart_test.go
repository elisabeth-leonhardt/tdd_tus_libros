package test

import (
	"TusLibros/source"
	"github.com/stretchr/testify/suite"
	"testing"
)

type CartTestSuite struct {
	suite.Suite
}

func (suite *CartTestSuite) TestNewCartIsEmpty() {
	aCart := EmptyCart()
	cantCart := len(aCart.Books)
	suite.Equal(0, cantCart)
}

func (suite *CartTestSuite) TestAddBookAndCartContainslt() {
	aCart := EmptyCart()
	aBook := ValidBook()
	aCart.AddObjectToCart(aBook, 1)
	cantCart := len(aCart.Books)
	suite.Equal(1, cantCart)
}

func (suite *CartTestSuite) TestAddTwoBooksAndCartContainsThem() {
	aCart := EmptyCart()
	aBook := ValidBook()
	anotherBook := ValidBook()
	aCart.AddObjectToCart(aBook, 1)
	aCart.AddObjectToCart(anotherBook, 1)
	cantCart := len(aCart.Books)
	suite.Equal(2, cantCart)
}

func (suite *CartTestSuite) TestAddTwoDifferentBooksAndCartContainsThem() {
	aCart := EmptyCart()
	aBook := ValidBook()
	aCart.AddObjectToCart(aBook, 3)
	cantCart := len(aCart.Books)
	suite.Equal(3, cantCart)
}

func (suite *CartTestSuite) TestCannotAdd0Books() {
	aCart := EmptyCart()
	aBook := ValidBook()
	err := aCart.AddObjectToCart(aBook, 0)
	cantCart := len(aCart.Books)
	suite.Equal(0, cantCart)
	suite.Equal(err, source.Error01)
}

func (suite *CartTestSuite) TestListBooksInCart() {
	aCart := EmptyCart()
	aBook := ValidBook()
	anotherBook := ValidBook()
	aCart.AddObjectToCart(aBook, 1)
	aCart.AddObjectToCart(anotherBook, 1)
	cartContent, _ := aCart.ListCartContent()
	var message = "[{\"isbn\":\"ValidBook\"},{\"isbn\":\"ValidBook\"}]"
	suite.Equal(message, cartContent)
}

func (suite *CartTestSuite) TestVerifyThatCatalogueContainsBook() {
	aCart := EmptyCart()
	aBookValid := ValidBook()
	err := aCart.AddObjectToCart(aBookValid, 1)
	suite.Equal(nil, err)
}

func (suite *CartTestSuite) TestVerifyThatCatalogueNotContainsBook() {
	aCart := EmptyCart()
	aBookNoValid := InvalidBook()
	err := aCart.AddObjectToCart(aBookNoValid, 1)
	suite.Equal(source.Error02, err)
}

func TestRunCartSuite(t *testing.T) {
	suite.Run(t, new(CartTestSuite))
}
