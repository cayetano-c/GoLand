package main

import "fmt"

func main() {
	notas := [6][4]float64{
		{9.5, 9.0, 9.0, 7.5},
		{5.0, 8.5, 7.5, 4.0},
		{9.5, 10.0, 7.5, 10.0},
		{4.5, 5.5, 6.5, 5.0},
		{8.0, 8.0, 7.5, 9.0},
		{6.0, 5.5, 6.5, 7.0},
	}
	totalClase := 0.0
	for i := 0; i < 6; i++ {
		fila := notas[i][:]
		suma := 0.0
		max, min := fila[0], fila[0]
		for _, n := range fila {
			suma += n
			if n > max {
				max = n
			}
			if n < min {
				min = n
			}
		}

		totalClase += suma
		fmt.Printf("Estudiante %d -> Promedio: %.2f | Máxima: %.1f | Mínima: %.1f\n", i+1, suma/4, max, min)
	}

	fmt.Printf("Promedio general de la clase: %.2f\n", totalClase/24)
}