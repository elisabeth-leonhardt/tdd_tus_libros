package test

import (
	"testing"
	"github.com/stretchr/testify/suite"

)

type CartTestSuite struct {
    suite.Suite
}

func (suite *CartTestSuite) Test001() {
    cart := 
}

func RunCartTestSuite(t *testing.T) {
    suite.Run(t, new(CartTestSuite))
}