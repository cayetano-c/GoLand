package main

import "fmt"

func ganador(votos map[string]int) string {
	mejor := ""
	max := -1
	for act, v := range votos {
		if v > max {
			max = v
			mejor = act
		}
	}
	return mejor
}

func main() {
	votos := map[string]int{"deportes": 0, "videojuegos": 0, "cine": 0, "musica": 0}
	opciones := []string{"deportes", "videojuegos", "cine", "musica"}

	fmt.Println("1. Deportes\n2. Videojuegos\n3. Cine\n4. Música")

	for i := 1; i <= 5; i++ {
		var n int
		fmt.Printf("Voto %d (1-4): ", i)
		fmt.Scan(&n)

		if n >= 1 && n <= 4 {
			votos[opciones[n-1]]++
		} else {
			fmt.Println("Opción no válida")
			i--
		}
	}

	for act, v := range votos {
		fmt.Println(act, ":", v)
	}

	fmt.Println("Ganador:", ganador(votos))
}
