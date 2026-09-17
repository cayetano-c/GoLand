package main

import "fmt"

/*
func <nombre>(param1, param2, ....param n)<valores de retorno>{
-----------------------------
-----------------------------
-----------------------------
// return en el caso de que nuestra función retorne valores
}


*/

func saludar() {
	fmt.Println("Hola esta es mi primera función")
}

func bienvenida(nombre string) {
	fmt.Println("Bienvenid@", nombre)
}

func main() {
	var usr string

	fmt.Println("Ingresa tu nombre: ")
	fmt.Scan(&usr)
	saludar()
	bienvenida(usr)

	suma, resta := sumaresta(10, 20)
	fmt.Println("El resultado de la suma y resta son: ", suma, " y ", resta)

	mostrarNum(5, 10, 15, 20, 25)
	fmt.Println("La sumatoria es: ", sumatoria(2, 4, 6, 8))
}

func sumaresta(a int, b int) (int, int) {
	if a < b {
		return b - a, a + b
	}
	return a + b, 0

}

/*Funcion Variadica*/
func mostrarNum(numeros ...int) {
	fmt.Println("Los numeros ingresados son: ", numeros)
}

func sumatoria(numeros ...int) int {
	total := 0
	for _, num := range numeros {
		total += num
	}
	return total
}
