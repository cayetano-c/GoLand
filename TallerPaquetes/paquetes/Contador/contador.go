package Contador

func ContarVocales(cadena string) (int, int, int, int, int) {
	var a, e, i, o, u int

	for _, letra := range cadena {
		switch letra {
		case 'a', 'A', 'á', 'Á':
			a++
		case 'e', 'E', 'é', 'É':
			e++
		case 'i', 'I', 'í', 'Í':
			i++
		case 'o', 'O', 'ó', 'Ó':
			o++
		case 'u', 'U', 'ú', 'Ú', 'ü', 'Ü':
			u++
		}
	}

	return a, e, i, o, u
}
