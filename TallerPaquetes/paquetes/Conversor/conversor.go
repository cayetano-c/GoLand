package Conversor

func DolarEuro(cantidadDolares float64) float64 {
	tipoCambio := 0.88
	cantidadEuros := cantidadDolares * tipoCambio
	return cantidadEuros
}

func DolarLibra(cantidadDolares float64) float64 {
	tipoCambio := 0.76
	cantidadLibras := cantidadDolares * tipoCambio
	return cantidadLibras
}

func DolarWon(cantidadDolares float64) float64 {
	tipoCambio := 1366.00
	cantidadWones := cantidadDolares * tipoCambio
	return cantidadWones
}

func DolarBTC(cantidadDolares float64) float64 {
	tipoCambio := 0.000012
	cantidadBTC := cantidadDolares * tipoCambio
	return cantidadBTC
}
