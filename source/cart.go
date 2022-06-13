package source

import "fmt"

type Cart struct {
	Id string
	Books []Book
}

type Book struct {
	ISBN string
	Author string
	Title string
}

func(c *Cart) GetCart() []Book {
	return c.Books
}

func(c *Cart) AddBook(aBook Book) {
	c.Books = append(c.Books, aBook)
	fmt.Print(c.Books)
}


// 1. primer carrito que agarro está vacío
// 2. agregar un libro y los contiene
// 3. agregar varios libros y los contiene
// 4. agregar 3x el mismo libro y los contiene
// 5. quiero agregar 0 o menos libros y no se puede (tira error)
// 6. 