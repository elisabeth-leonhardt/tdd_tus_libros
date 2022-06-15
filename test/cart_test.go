package test

import (
	"TusLibros/source"
	"github.com/stretchr/testify/suite"
	"testing"
)

type CartTestSuite struct {
	suite.Suite
	factory Factory
}

func (suite *CartTestSuite) SetupTest() {
	suite.factory = NewFactory()
}

func (suite *CartTestSuite) TestNewCartIsEmpty() {
	aCart := suite.factory.EmptyCart()
	cantCart := len(aCart.Books)
	suite.Equal(0, cantCart)
}

func (suite *CartTestSuite) TestAddBookAndCartContainslt() {
	aCart := suite.factory.EmptyCart()
	aBook := suite.factory.ValidBook()
	aCart.AddObjectToCart(aBook, 1)
	cantCart := len(aCart.Books)
	suite.Equal(1, cantCart)
}

func (suite *CartTestSuite) TestAddTwoBooksAndCartContainsThem() {
	aCart := suite.factory.EmptyCart()
	aBook := suite.factory.ValidBook()
	anotherBook := suite.factory.ValidBook()
	aCart.AddObjectToCart(aBook, 1)
	aCart.AddObjectToCart(anotherBook, 1)
	cantCart := len(aCart.Books)
	suite.Equal(2, cantCart)
}

func (suite *CartTestSuite) TestAddTwoDifferentBooksAndCartContainsThem() {
	aCart := suite.factory.EmptyCart()
	aBook := suite.factory.ValidBook()
	aCart.AddObjectToCart(aBook, 3)
	cantCart := len(aCart.Books)
	suite.Equal(3, cantCart)
}

func (suite *CartTestSuite) TestCannotAdd0Books() {
	aCart := suite.factory.EmptyCart()
	aBook := suite.factory.ValidBook()
	err := aCart.AddObjectToCart(aBook, 0)
	cantCart := len(aCart.Books)
	suite.Equal(0, cantCart)
	suite.Equal(err, source.Error01)
}

func (suite *CartTestSuite) TestListBooksInCart() {
	aCart := suite.factory.EmptyCart()
	aBook := suite.factory.ValidBook()
	anotherBook := suite.factory.ValidBook()
	aCart.AddObjectToCart(aBook, 1)
	aCart.AddObjectToCart(anotherBook, 1)
	cartContent, _ := aCart.ListCartContent()
	var message = "[{\"isbn\":\"ValidBook\"},{\"isbn\":\"ValidBook\"}]"
	suite.Equal(message, cartContent)
}

func (suite *CartTestSuite) TestVerifyThatCatalogueContainsBook() {
	aCart := suite.factory.EmptyCart()
	aBookValid := suite.factory.ValidBook()
	err := aCart.AddObjectToCart(aBookValid, 1)
	suite.Equal(nil, err)
}

func (suite *CartTestSuite) TestVerifyThatCatalogueNotContainsBook() {
	aCart := suite.factory.EmptyCart()
	aBookNoValid := suite.factory.InvalidBook()
	err := aCart.AddObjectToCart(aBookNoValid, 1)
	suite.Equal(source.Error02, err)
}

func TestRunCartSuite(t *testing.T) {
	suite.Run(t, new(CartTestSuite))
}
