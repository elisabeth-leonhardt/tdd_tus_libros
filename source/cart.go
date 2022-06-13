package source

import (
	"encoding/json"
	"errors"
)

type Cart struct {
	Books     []Book
	Catalogue []Book
}

type Book struct {
	ISBN string `json:"isbn"`
}

var Error01 = errors.New("cannot receive 1 items")
var Error02 = errors.New("book is not part of catalogue")

// 1. primer carrito que agarro está vacío
// 2. agregar un libro y los contiene
// 3. agregar varios libros y los contiene
// 4. agregar 3x el mismo libro y los contiene
// 5. quiero agregar 0 o menos libros y no se puede (tira error)

func NewCart(catalogueBooks []Book) Cart {
	newCart := Cart{Books: []Book{}, Catalogue: catalogueBooks}
	return newCart
}

func (c *Cart) AddObjectToCart(aBook Book, numberOfCopies int) (err error) {
	if numberOfCopies == 0 {
		return Error01
	}

	flag := false
	for _, b := range c.Catalogue {
		if aBook.ISBN == b.ISBN {
			for i := 0; i < numberOfCopies; i++ {
				c.Books = append(c.Books, aBook)
			}
			flag = true
		}
	}

	if flag == false {
		return Error02
	}

	return nil
}

func (c *Cart) ListCartContent() (cartContent string, err error) {
	cartContentMarshal, err := json.Marshal(c.Books)
	cartContent = string(cartContentMarshal)
	return cartContent, err
}
