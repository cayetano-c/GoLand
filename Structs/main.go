package main

import "fmt"

type Producto struct {
	Nombre string
	precio float64
	stock  int
}

func NewProduct(nombre string, precio float64, stock int) Producto {
	return Producto{
		Nombre: nombre,
		precio: precio,
		stock:  stock,
	}
}

func CalcularValorInventario(p Producto) float64 {
	return p.precio * float64(p.stock)

}
func (p *Producto) Vender(cantidad int) bool {
	if cantidad < 0 || cantidad > p.stock {
		return false
	}
	p.stock -= cantidad
	return true
}
func main() {

	producto := Producto{}
	producto.Nombre = "Intel Core Ultra 7"
	producto.precio = 850.90
	producto.stock = 200

	producto2 := Producto{
		Nombre: "Teclado",
		precio: 30.5,
		stock:  15,
	}

	fmt.Println(producto)
	fmt.Println(producto2)
	fmt.Println("El Producto: ", producto2.Nombre, "\nValor del inventario: ", CalcularValorInventario(producto2))

	product3 := NewProduct("mouse", 12.91, 50)
	fmt.Println(product3)
}
