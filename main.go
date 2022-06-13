package main

import (
	"TusLibros/source"
	"fmt"
)

func main() {
	aCart := source.Cart{}
	fmt.Print(aCart)
	aCart.AddBook("titulo de un libro")
	fmt.Print(aCart)
}