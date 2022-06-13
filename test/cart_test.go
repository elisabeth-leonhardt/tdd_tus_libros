package test

import (
	"TusLibros/source"
	"testing"

	"github.com/stretchr/testify/suite"
)

type CartTestSuite struct {
	suite.Suite
}

//carrito esta vacio
func (suite *CartTestSuite) Test001() {
    aCart := suite.emptyCart()
	cantCart := len(aCart.Books)
	suite.Equal(0, cantCart)
}

//agregar un libro y este lo contiene
func (suite *CartTestSuite) Test002() {
    aCart := suite.emptyCart()
    aBook := suite.validBook()
	aCart.AddObjectToCart(aBook, 1)
	cantCart := len(aCart.Books)
	suite.Equal(1, cantCart)
}

//agregar varios libros y este lo contiene
func (suite *CartTestSuite) Test003() {
    aCart := suite.emptyCart()
    aBook := suite.validBook()
    anotherBook := suite.validBook()
    aCart.AddObjectToCart(aBook, 1)
    aCart.AddObjectToCart(anotherBook, 1)
	cantCart := len(aCart.Books)
	suite.Equal(2, cantCart)
}






func (suite *CartTestSuite) Test004() {
    aCart := suite.emptyCart()
    aBook := suite.validBook()
    aCart.AddObjectToCart(aBook, 3)
	cantCart := len(aCart.Books)
	suite.Equal(3, cantCart)
}

func (suite *CartTestSuite) Test005() {
    aCart := suite.emptyCart()
    aBook := suite.validBook()
    err := aCart.AddObjectToCart(aBook, 0)
	cantCart := len(aCart.Books)
    suite.Equal(0, cantCart)
    suite.Equal(err, source.Error01)
}

// listar los libros
func (suite *CartTestSuite) Test006() {
    aCart := suite.emptyCart()
    aBook := suite.validBook()
    anotherBook := suite.validBook()
    aCart.AddObjectToCart(aBook, 1)
    aCart.AddObjectToCart(anotherBook, 1)
    cartContent, _ := aCart.ListCartContent()
    var message = "{}"
    suite.Equal(message, cartContent)

}


// Factory Methods
func (*CartTestSuite) validBook() (source.Book) {
	return source.Book{ISBN: "validBook"}
}

func (*CartTestSuite) emptyCart() (source.Cart) {
	return source.NewCart()
}

//Esta funcion tiene que comenzar con Test
func TestRunCartSuite(t *testing.T) {
	suite.Run(t, new(CartTestSuite))
}
