package test

import (
	"TusLibros/source"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type CartTestSuite struct {
    suite.Suite
}

func (suite *CartTestSuite) TestExample001() {
    aCart := source.Cart{}
    suite.Equal(len(aCart.GetCart()), 0)
}

func (suite *CartTestSuite) Test002() {
    aCart := source.Cart{}
    aBook := source.Book{Author: "Cal Newport", Title: "Deep Work", ISBN: "123"}
    aCart.AddBook(aBook)
    suite.Equal(len(aCart.GetCart()), 1)
}

// this works!!
func Test001(t *testing.T) {
    aCart := source.Cart{}
	assert.Equal(t, 0, len(aCart.GetCart()));
}

func Test002(t *testing.T) {
    aCart := source.Cart{}
    aBook := source.Book{Author: "Cal Newport", Title: "Deep Work", ISBN: "123"}
    aCart.AddBook(aBook)
	assert.Equal(t, 1, len(aCart.GetCart()));
}

func RunCartTestSuite(t *testing.T) {
    suite.Run(t, new(CartTestSuite))
}