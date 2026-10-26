package main

import "fmt"

type Producto struct {
	ID       int
	Nombre   string
	Precio   float64
	Cantidad int
}

func NuevoProducto(id int, nombre string, precio float64, cantidad int) Producto {
	return Producto{
		ID:       id,
		Nombre:   nombre,
		Precio:   precio,
		Cantidad: cantidad,
	}
}
func (p *Producto) ActualizarPrecio(nuevoPrecio float64) {
	p.Precio = nuevoPrecio
}

func (p *Producto) ActualizarCantidad(nuevaCantidad int) {
	p.Cantidad = nuevaCantidad
}

func (p Producto) MostrarInfo() {
	fmt.Printf("ID:       %d\n", p.ID)
	fmt.Printf("Nombre:   %s\n", p.Nombre)
	fmt.Printf("Precio:   $%.2f\n", p.Precio)
	fmt.Printf("Cantidad: %d unidades\n", p.Cantidad)
}

type Inventario struct {
	Productos map[int]Producto
}

func NuevoInventario() Inventario {
	return Inventario{
		Productos: make(map[int]Producto),
	}
}
func (inv *Inventario) AgregarProducto(p Producto) {
	if _, existe := inv.Productos[p.ID]; existe {
		fmt.Printf("Error: ya existe un producto con ID %d\n", p.ID)
		return
	}
	inv.Productos[p.ID] = p
	fmt.Printf("Producto '%s' agregado correctamente.\n", p.Nombre)
}
func (inv *Inventario) EliminarProducto(id int) {
	if _, existe := inv.Productos[id]; !existe {
		fmt.Printf("Error: no existe un producto con ID %d\n", id)
		return
	}
	delete(inv.Productos, id)
	fmt.Printf("Producto con ID %d eliminado.\n", id)
}
func (inv Inventario) BuscarProducto(id int) (Producto, bool) {
	p, existe := inv.Productos[id]
	return p, existe
}
func (inv Inventario) ListarProductos() {
	if len(inv.Productos) == 0 {
		fmt.Println("El inventario está vacío.")
		return
	}
	fmt.Println("===== INVENTARIO =====")
	for _, p := range inv.Productos {
		p.MostrarInfo()
	}
}
func main() {
	inventario := NuevoInventario()
	p1 := NuevoProducto(1, "Laptop", 850.50, 10)
	p2 := NuevoProducto(2, "Mouse", 15.99, 50)
	p3 := NuevoProducto(3, "Teclado", 45.00, 30)
	inventario.AgregarProducto(p1)
	inventario.AgregarProducto(p2)
	inventario.AgregarProducto(p3)
	inventario.AgregarProducto(NuevoProducto(1, "Monitor", 200, 5))
	fmt.Println()
	inventario.ListarProductos()
	fmt.Println("\nBuscando el producto con ID 2...")
	if p, encontrado := inventario.BuscarProducto(2); encontrado {
		p.MostrarInfo()
		p.ActualizarPrecio(12.50)
		p.ActualizarCantidad(45)
		inventario.Productos[p.ID] = p
		fmt.Println("Producto actualizado:")
		p.MostrarInfo()
	} else {
		fmt.Println("Producto no encontrado.")
	}
	if _, encontrado := inventario.BuscarProducto(99); !encontrado {
		fmt.Println("El producto con ID 99 no existe.")
	}
	fmt.Println()
	inventario.EliminarProducto(3)
	inventario.EliminarProducto(3)
	fmt.Println()
	inventario.ListarProductos()
}
