package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	lector := bufio.NewReader(os.Stdin)
	fmt.Print("Ingrese una frase o texto: ")
	texto, _ := lector.ReadString('\n')

	texto = strings.ToLower(texto)

	signos := []string{",", ".", ";", ":", "!", "?", "¿", "¡"}
	for _, s := range signos {
		texto = strings.ReplaceAll(texto, s, "")
	}

	palabras := strings.Fields(texto)

	conteo := make(map[string]int)
	for _, palabra := range palabras {
		conteo[palabra]++
	}

	fmt.Println("\nConteo de palabras:")
	for palabra, veces := range conteo {
		fmt.Printf("%s: %d\n", palabra, veces)
	}

	masRepetida := ""
	maximo := 0
	for palabra, veces := range conteo {
		if veces > maximo {
			maximo = veces
			masRepetida = palabra
		}
	}

	if maximo > 0 {
		fmt.Printf("\nLa palabra que más se repite es \"%s\" (%d veces)\n", masRepetida, maximo)
	} else {
		fmt.Println("\nNo se ingresó ninguna palabra.")
	}
}
