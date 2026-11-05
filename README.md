# GoLand
package main

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

const dateLayout = "2006-01-02"

// ---------------------------------------------------------------------------
// Medicine (estructura base)
// ---------------------------------------------------------------------------

type Medicine struct {
	name             string
	manufacturer     string
	manufactureDate  time.Time
	shelfLife        int // vida útil en meses
}

// NewMedicine es el constructor base: valida todos los campos.
// La fecha se recibe como texto "AAAA-MM-DD"; time.Parse rechaza fechas inexistentes (p. ej. 2024-02-30).
func NewMedicine(name, manufacturer, manufactureDate string, shelfLife int) (*Medicine, error) {
	m := &Medicine{}
	if err := m.SetName(name); err != nil {
		return nil, err
	}
	if err := m.SetManufacturer(manufacturer); err != nil {
		return nil, err
	}
	if err := m.SetManufactureDate(manufactureDate); err != nil {
		return nil, err
	}
	if err := m.SetShelfLife(shelfLife); err != nil {
		return nil, err
	}
	return m, nil
}

// Getters
func (m *Medicine) Name() string            { return m.name }
func (m *Medicine) Manufacturer() string    { return m.manufacturer }
func (m *Medicine) ManufactureDate() time.Time { return m.manufactureDate }
func (m *Medicine) ShelfLife() int          { return m.shelfLife }

// Setters con validación
func (m *Medicine) SetName(name string) error {
	if strings.TrimSpace(name) == "" {
		return errors.New("el nombre no puede estar vacío")
	}
	m.name = strings.TrimSpace(name)
	return nil
}

func (m *Medicine) SetManufacturer(manufacturer string) error {
	if strings.TrimSpace(manufacturer) == "" {
		return errors.New("el fabricante no puede estar vacío")
	}
	m.manufacturer = strings.TrimSpace(manufacturer)
	return nil
}

func (m *Medicine) SetManufactureDate(date string) error {
	parsed, err := time.Parse(dateLayout, strings.TrimSpace(date))
	if err != nil {
		return fmt.Errorf("fecha de fabricación inválida %q (formato esperado AAAA-MM-DD): %w", date, err)
	}
	m.manufactureDate = parsed
	return nil
}

func (m *Medicine) SetShelfLife(months int) error {
	if months <= 0 {
		return errors.New("la vida útil debe ser un número positivo de meses")
	}
	m.shelfLife = months
	return nil
}

// ExpirationDate calcula la fecha de caducidad: fabricación + vida útil (meses).
func (m *Medicine) ExpirationDate() time.Time {
	return m.manufactureDate.AddDate(0, m.shelfLife, 0)
}

// ---------------------------------------------------------------------------
// Tablet ("hereda" de Medicine mediante composición/embedding)
// ---------------------------------------------------------------------------

type Tablet struct {
	Medicine
	dosePerTablet  float64 // mg
	isPrescription bool
}

func NewTablet(name, manufacturer, manufactureDate string, shelfLife int,
	dosePerTablet float64, isPrescription bool) (*Tablet, error) {

	base, err := NewMedicine(name, manufacturer, manufactureDate, shelfLife)
	if err != nil {
		return nil, err
	}
	t := &Tablet{Medicine: *base, isPrescription: isPrescription}
	if err := t.SetDosePerTablet(dosePerTablet); err != nil {
		return nil, err
	}
	return t, nil
}

func (t *Tablet) DosePerTablet() float64 { return t.dosePerTablet }
func (t *Tablet) IsPrescription() bool   { return t.isPrescription }

func (t *Tablet) SetDosePerTablet(dose float64) error {
	if dose <= 0 {
		return errors.New("la dosis por tableta debe ser mayor a 0")
	}
	t.dosePerTablet = dose
	return nil
}

func (t *Tablet) SetIsPrescription(v bool) { t.isPrescription = v }

func (t *Tablet) ShowDetails() {
	fmt.Println("--- Tableta ---")
	fmt.Printf("Nombre:             %s\n", t.Name())
	fmt.Printf("Fabricante:         %s\n", t.Manufacturer())
	fmt.Printf("Fecha fabricación:  %s\n", t.ManufactureDate().Format(dateLayout))
	fmt.Printf("Vida útil:          %d meses\n", t.ShelfLife())
	fmt.Printf("Dosis por tableta:  %.1f mg\n", t.DosePerTablet())
	fmt.Printf("Requiere receta:    %t\n", t.IsPrescription())
}

// ---------------------------------------------------------------------------
// Syrup
// ---------------------------------------------------------------------------

type Syrup struct {
	Medicine
	volume float64 // ml
	flavor string
}

func NewSyrup(name, manufacturer, manufactureDate string, shelfLife int,
	volume float64, flavor string) (*Syrup, error) {

	base, err := NewMedicine(name, manufacturer, manufactureDate, shelfLife)
	if err != nil {
		return nil, err
	}
	s := &Syrup{Medicine: *base, flavor: flavor}
	if err := s.SetVolume(volume); err != nil {
		return nil, err
	}
	return s, nil
}

func (s *Syrup) Volume() float64 { return s.volume }
func (s *Syrup) Flavor() string  { return s.flavor }

func (s *Syrup) SetVolume(volume float64) error {
	if volume <= 0 {
		return errors.New("el volumen debe ser mayor a 0")
	}
	s.volume = volume
	return nil
}

func (s *Syrup) SetFlavor(flavor string) { s.flavor = flavor }

func (s *Syrup) ShowDetails() {
	fmt.Println("--- Jarabe ---")
	fmt.Printf("Nombre:             %s\n", s.Name())
	fmt.Printf("Fabricante:         %s\n", s.Manufacturer())
	fmt.Printf("Fecha fabricación:  %s\n", s.ManufactureDate().Format(dateLayout))
	fmt.Printf("Vida útil:          %d meses\n", s.ShelfLife())
	fmt.Printf("Volumen:            %.0f ml\n", s.Volume())
	fmt.Printf("Sabor:              %s\n", s.Flavor())
}

// ---------------------------------------------------------------------------
// main
// ---------------------------------------------------------------------------

// must evita repetir el manejo de errores al construir los datos de ejemplo.
func must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}

func main() {
	// 1) Dos listas de inventario: una de Tablet y otra de Syrup
	tablets := []*Tablet{
		must(NewTablet("Paracetamol", "Genfar", "2024-03-15", 24, 500, false)),
		must(NewTablet("Amoxicilina", "Pfizer", "2024-08-01", 18, 875, true)),
	}
	syrups := []*Syrup{
		must(NewSyrup("Ambroxol", "Roemmers", "2025-01-10", 12, 120, "Fresa")),
		must(NewSyrup("Ibuprofeno", "Bayer", "2024-11-20", 36, 100, "Naranja")),
	}

	// 2) Mostrar detalles de todos los medicamentos
	fmt.Println("===== INVENTARIO DE TABLETAS =====")
	for _, t := range tablets {
		t.ShowDetails()
		fmt.Println()
	}
	fmt.Println("===== INVENTARIO DE JARABES =====")
	for _, s := range syrups {
		s.ShowDetails()
		fmt.Println()
	}

	// 3) Actualizar atributos con setters y mostrar el cambio con getters
	fmt.Println("===== ACTUALIZACIONES =====")
	fmt.Printf("Dosis anterior de %s: %.1f mg\n", tablets[0].Name(), tablets[0].DosePerTablet())
	if err := tablets[0].SetDosePerTablet(650); err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Printf("Dosis nueva de %s:    %.1f mg\n", tablets[0].Name(), tablets[0].DosePerTablet())

	// Intentos inválidos: los setters devuelven error y no modifican el valor
	if err := syrups[0].SetVolume(-5); err != nil {
		fmt.Println("Error esperado:", err)
	}
	if err := syrups[1].SetName("   "); err != nil {
		fmt.Println("Error esperado:", err)
	}
	if err := tablets[1].SetManufactureDate("2024-02-30"); err != nil {
		fmt.Println("Error esperado:", err)
	}
	if err := tablets[1].SetShelfLife(0); err != nil {
		fmt.Println("Error esperado:", err)
	}
	fmt.Println()

	// 4) Fecha de caducidad de cada medicamento
	fmt.Println("===== FECHAS DE CADUCIDAD =====")
	for _, t := range tablets {
		fmt.Printf("%-12s (tableta) caduca el %s\n", t.Name(), t.ExpirationDate().Format(dateLayout))
	}
	for _, s := range syrups {
		fmt.Printf("%-12s (jarabe)  caduca el %s\n", s.Name(), s.ExpirationDate().Format(dateLayout))
	}
}