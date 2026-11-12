package main

import (
	"errors"
	"fmt"
)

type Medicine struct {
	name            string
	manufacturer    string
	manufactureDate string
	shelfLife       int
}

func (m Medicine) GetName() string {
	return m.name
}
func (m *Medicine) SetName(name string) error {
	if name == "" {
		return errors.New("Error: El texto no puede estar vacío. Coloque un nombre: ")
	}
	m.name = name
	return nil
}
func (m Medicine) GetManufactureDate() string {
	return m.manufactureDate
}
func (m *Medicine) SetManufactureDate(manufactureDate string) error {
	if manufactureDate == "" {
		return errors.New("Error: El texto no puede estar vacío. Coloque una fecha: ")
	}
	m.manufactureDate = manufactureDate
	return nil
}
func (m Medicine) GetShelfLife() int {
	return m.shelfLife
}
func (m *Medicine) SetShelfLife(shelfLife int) error {
	if shelfLife <= 0 {
		return errors.New("Error: La vida útil debe ser un número positivo")
	}
	m.shelfLife = shelfLife
	return nil
}

type Tablet struct {
	Medicine
	dosePerTablet  float64
	isPrescription bool
}
type Syrup struct {
	Medicine
	volume float64
	flavor string
}

func (m Medicine) GetManufacturer() string {
	return m.manufacturer
}
func (m *Medicine) SetManufacturer(manufacturer string) error {
	if manufacturer == "" {
		return errors.New("Error: El fabricante no puede estar vacío.")
	}
	m.manufacturer = manufacturer
	return nil
}
func (m Medicine) MostrarCaducidad() {
	fmt.Println("Fabricado en:", m.manufactureDate, "| Vida útil:", m.shelfLife, "meses")
}
func (t Tablet) MostrarDetalles() {
	fmt.Println("--- TABLETA ---")
	fmt.Println("Nombre:", t.GetName())
	fmt.Println("Fabricante:", t.GetManufacturer())
	fmt.Println("Fecha de Fabricación:", t.GetManufactureDate())
	fmt.Println("Dosis:", t.dosePerTablet, "mg")
	fmt.Println("Requiere receta:", t.isPrescription)
	t.MostrarCaducidad()
	fmt.Println()
}
func (s Syrup) MostrarDetalles() {
	fmt.Println("--- JARABE ---")
	fmt.Println("Nombre:", s.GetName())
	fmt.Println("Fabricante:", s.GetManufacturer())
	fmt.Println("Fecha de Fabricación:", s.GetManufactureDate())
	fmt.Println("Volumen:", s.volume, "ml")
	fmt.Println("Sabor:", s.flavor)
	s.MostrarCaducidad()
	fmt.Println()
}
func main() {
	var listaTabletas []Tablet
	var listaJarabes []Syrup

	var t1 Tablet
	t1.name = "Paracetamol"
	t1.manufacturer = "Bayer"
	t1.manufactureDate = "20-05-2025"
	t1.shelfLife = 24
	t1.dosePerTablet = 500.0
	t1.isPrescription = false

	var s1 Syrup
	s1.name = "Jarabe"
	s1.manufacturer = "Pfizer"
	s1.manufactureDate = "15-11-2025"
	s1.shelfLife = 12
	s1.volume = 120.0
	s1.flavor = "Miel"
	listaTabletas = append(listaTabletas, t1)
	listaJarabes = append(listaJarabes, s1)
}
