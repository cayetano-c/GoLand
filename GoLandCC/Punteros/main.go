package main

import "fmt"

func main() {
	numero := 10

	puntero := &numero

	fmt.Println("Valor: ", numero)
	fmt.Println("Dirección de memoria: ", &numero)
	fmt.Println("Puntero: ", puntero)
	fmt.Println("Puntero: ", *puntero)
}
