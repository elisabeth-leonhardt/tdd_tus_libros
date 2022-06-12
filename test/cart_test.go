package test

import (
	"TusLibros/source"
	"github.com/stretchr/testify/suite"
	"testing"
)

type CartTestSuite struct {
	suite.Suite
}

//carrito esta vacio
func (suite *CartTestSuite) Test001() {
	cart := source.GetCart()
	cantCart := len(cart)
	suite.Equal(0, cantCart)
}

//agregar un libro y este lo contiene
func (suite *CartTestSuite) Test002() {
	idsBooks := []int{001}
	cart := source.AddObjectToCart(idsBooks)
	cantCart := len(cart)
	suite.Equal(1, cantCart)
}

//agregar varios libros y este lo contiene
func (suite *CartTestSuite) Test003() {
	idsBooks := []int{001, 002}
	cart := source.AddObjectToCart(idsBooks)
	cantCart := len(cart)
	suite.Equal(2, cantCart)
}

//Esta funcion tiene que comenzar con Test.. por eso no te lo tomaba
func TestRunCartSuite(t *testing.T) {
	suite.Run(t, new(CartTestSuite))
}
