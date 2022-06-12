package source

type Cart struct {
	quant int
}

// 1. primer carrito que agarro está vacío
// 2. agregar un libro y los contiene
// 3. agregar varios libros y los contiene
// 4. agregar 3x el mismo libro y los contiene
// 5. quiero agregar 0 o menos libros y no se puede (tira error)
// 6.

/*func CreateCart() Cart {
	newCart := Cart{quant: 0}
	return newCart
}*/

func GetCart() (cart []int) {
	return
}

func AddObjectToCart(idBook []int) (cart []int) {
	cart = append(cart, idBook...)
	return cart
}
