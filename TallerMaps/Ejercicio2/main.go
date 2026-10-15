package main

import "fmt"

type Producto struct {
	Nombre   string
	Precio   float64
	Cantidad int
}

func mostrarProducto(p Producto) {
	fmt.Printf("Nombre: %s\n", p.Nombre)
	fmt.Printf("Precio: $%.2f\n", p.Precio)
	fmt.Printf("Cantidad: %d\n", p.Cantidad)
}

func valorTotal(p Producto) float64 {
	return p.Precio * float64(p.Cantidad)
}

func tieneStock(p Producto) bool {
	return p.Cantidad > 0
}

func main() {
	productos := []Producto{
		{Nombre: "Laptop", Precio: 850.50},
		{Nombre: "Mouse", Precio: 15.99},
		{Nombre: "Teclado", Precio: 45.00},
	}

	for i := range productos {
		fmt.Printf("Ingrese el stock de %s: ", productos[i].Nombre)
		fmt.Scan(&productos[i].Cantidad)
	}

	fmt.Println("\n=== Lista de productos ===")
	for _, p := range productos {
		mostrarProducto(p)
		fmt.Printf("Valor total: $%.2f\n", valorTotal(p))

		if tieneStock(p) {
			fmt.Println("Estado de stock: disponible")
		} else {
			fmt.Println("Estado de stock: sin stock")
		}
		fmt.Println("---------------------------")
	}
}
