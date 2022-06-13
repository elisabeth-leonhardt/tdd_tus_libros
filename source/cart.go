package source

import (
	"encoding/json"
	"errors"
) 
type Cart struct {
	Books []Book
}

type Book struct {
	ISBN string `json:"isbn"`
}

var Error01 = errors.New("cannot receive 1 items")

// 1. primer carrito que agarro está vacío
// 2. agregar un libro y los contiene
// 3. agregar varios libros y los contiene
// 4. agregar 3x el mismo libro y los contiene
// 5. quiero agregar 0 o menos libros y no se puede (tira error)

func NewCart() Cart {
	newCart := Cart{Books: []Book{}}
	return newCart
}

func GetCart() (cart []int) {
	return
}

func (c *Cart) AddObjectToCart(aBook Book, numberOfCopies int) (err error) {
	if(numberOfCopies == 0) {
		return Error01
	}
	for i := 0; i< numberOfCopies; i++ {
		c.Books = append(c.Books, aBook)
	}
	return nil
}

func (c *Cart) ListCartContent() (cartContent string, err error) {

	cartContent, err = json.Marshal(c.Books)
	return cartContent, err

}

// agregar catálogo!