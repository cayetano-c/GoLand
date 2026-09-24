package operaciones

func Suma(a, b int) int {
	return a + b
}

func Sumaresta(a int, b int) (int, int) {
	if a < b {
		return b - a, a + b
	}
	return a + b, 0

}

func Sumatoria(numeros ...int) int {
	total := 0
	for _, num := range numeros {
		total += num
	}
	return total
}
