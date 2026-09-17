package main

import "fmt"

// Opción 1: Promedio de notas
func averageGrade(cantidad int) float64 {
	var suma float64 = 0
	var nota float64

	for i := 1; i <= cantidad; i++ {
		fmt.Printf("Ingrese la nota del estudiante %d (0-100): ", i)
		fmt.Scan(&nota)
		suma = suma + nota
	}

	promedio := suma / float64(cantidad)
	return promedio
}

func opcionNotas() {
	var cantidad int
	fmt.Print("Ingrese la cantidad de estudiantes: ")
	fmt.Scan(&cantidad)

	promedio := averageGrade(cantidad)
	fmt.Printf("El promedio del curso es: %.2f\n", promedio)

	if promedio >= 70 {
		fmt.Println("El curso está APROBADO")
	} else {
		fmt.Println("El curso está REPROBADO")
	}

	switch {
	case promedio >= 90:
		fmt.Println("Excellent performance")
	case promedio >= 80:
		fmt.Println("Good performance")
	case promedio >= 70:
		fmt.Println("Satisfactory performance")
	default:
		fmt.Println("Needs improvement")
	}
}

// Opción 2: Suma de números del 1 al n
func sumaHastaN(n int) int {
	suma := 0
	for i := 1; i <= n; i++ {
		suma = suma + i
	}
	return suma
}

func opcionSuma() {
	var n int
	fmt.Print("Ingrese el número n: ")
	fmt.Scan(&n)

	resultado := sumaHastaN(n)
	fmt.Printf("La suma de 1 al %d es: %d\n", n, resultado)
}

// Opción 3: Celsius a Fahrenheit
func celsiusAFahrenheit(celsius float64) float64 {
	return (celsius * 9 / 5) + 32
}

func opcionCelsius() {
	var celsius float64
	fmt.Print("Ingrese la temperatura en Celsius: ")
	fmt.Scan(&celsius)

	fahrenheit := celsiusAFahrenheit(celsius)
	fmt.Printf("%.2f°C equivalen a %.2f°F\n", celsius, fahrenheit)
}

// Opción 4: Fahrenheit a Celsius
func fahrenheitACelsius(fahrenheit float64) float64 {
	return (fahrenheit - 32) * 5 / 9
}

func opcionFahrenheit() {
	var fahrenheit float64
	fmt.Print("Ingrese la temperatura en Fahrenheit: ")
	fmt.Scan(&fahrenheit)

	celsius := fahrenheitACelsius(fahrenheit)
	fmt.Printf("%.2f°F equivalen a %.2f°C\n", fahrenheit, celsius)
}

// Menú principal
func mostrarMenu() {
	fmt.Println("\n===== MENÚ =====")
	fmt.Println("1. Calcular promedio de notas")
	fmt.Println("2. Suma de números del 1 al n")
	fmt.Println("3. Convertir Celsius a Fahrenheit")
	fmt.Println("4. Convertir Fahrenheit a Celsius")
	fmt.Println("0. Salir (o escriba 'salir')")
	fmt.Print("Elija una opción: ")
}

func main() {
	var opcion string

	for {
		mostrarMenu()
		fmt.Scan(&opcion)

		switch opcion {
		case "1":
			opcionNotas()
		case "2":
			opcionSuma()
		case "3":
			opcionCelsius()
		case "4":
			opcionFahrenheit()
		case "0", "salir":
			fmt.Println("¡Hasta luego!")
			return
		default:
			fmt.Println("Opción no válida, intente de nuevo.")
		}
	}
}
