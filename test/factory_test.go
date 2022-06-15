package test

import (
	"TusLibros/source"
)

// Factory Methods
func ValidBook() source.Book {
	return source.Book{ISBN: "ValidBook"}
}

func InvalidBook() source.Book {
	return source.Book{ISBN: "InvalidBook"}
}

func EmptyCart() source.Cart {
	bookValid := source.Book{ISBN: "ValidBook"}
	booksCatalogue := []source.Book{bookValid}
	return source.NewCart(booksCatalogue)
}
