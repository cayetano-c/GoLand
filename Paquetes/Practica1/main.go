package main

import (
	"Practica/operaciones"
	"Practica/saludo"
	"fmt"
)

func main() {
	fmt.Println("----Bienvenid@s a la Clase de Paquetes----")
	mensaje := saludo.Saludar("Daral")
	fmt.Println(mensaje)

	suma := operaciones.Suma(5, 10)
	fmt.Println("La suma es: ", suma)
	suma, resta := operaciones.Sumaresta(10, 20)
	fmt.Println("El resultado de la suma y resta son: ", suma, resta)
	sumatoria := operaciones.Sumatoria(2, 4, 6, 8)
	fmt.Println("La sumatoria es: ", sumatoria)

}
