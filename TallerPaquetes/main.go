package main

import (
	"Taller/paquetes/Contador"
	"Taller/paquetes/Conversor"
	"bufio"
	"fmt"
	"os"
)

func main() {
	var opcion string

	for {
		fmt.Println("\n==============================")
		fmt.Println("             MENÚ")
		fmt.Println("==============================")
		fmt.Println("1. Conversor de monedas")
		fmt.Println("2. Contador de vocales")
		fmt.Println("3. Salir")
		fmt.Print("Seleccione una opción: ")

		fmt.Scan(&opcion)

		if opcion == "1" {
			var dolares float64
			fmt.Print("\nIngrese la cantidad en Dólares (USD): ")
			fmt.Scan(&dolares)

			fmt.Println("\nSeleccione la moneda destino:")
			fmt.Println("1. Euros")
			fmt.Println("2. Libras Esterlinas (LB)")
			fmt.Println("3. Won Surcoreano (Won)")
			fmt.Println("4. Bitcoin (BTC)")

			var subOpcion string
			fmt.Print("Elija una opción (1-4): ")
			fmt.Scan(&subOpcion)

			if subOpcion == "1" {
				res := Conversor.DolarEuro(dolares)
				fmt.Printf("Resultado: %.2f Euros\n", res)
			} else if subOpcion == "2" {
				res := Conversor.DolarLibra(dolares)
				fmt.Printf("Resultado: %.2f Libras Esterlinas (LB)\n", res)
			} else if subOpcion == "3" {
				res := Conversor.DolarWon(dolares)
				fmt.Printf("Resultado: %.2f Wones\n", res)
			} else if subOpcion == "4" {
				res := Conversor.DolarBTC(dolares)
				fmt.Printf("Resultado: %.8f BTC\n", res)
			} else {
				fmt.Println("Opción de moneda no válida.")
			}

		} else if opcion == "2" {
			fmt.Print("\nEscriba una palabra o frase: ")

			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			frase := scanner.Text()

			a, e, i, o, u := Contador.ContarVocales(frase)

			fmt.Println("\nConteo de vocales:")
			fmt.Printf("Vocal 'a': %d\n", a)
			fmt.Printf("Vocal 'e': %d\n", e)
			fmt.Printf("Vocal 'i': %d\n", i)
			fmt.Printf("Vocal 'o': %d\n", o)
			fmt.Printf("Vocal 'u': %d\n", u)

		} else if opcion == "3" {
			fmt.Println("\n¡Gracias por usar el programa! Hasta luego.")
			return

		} else {
			fmt.Println("\nOpción inválida. Intente de nuevo.")
		}
	}
}
