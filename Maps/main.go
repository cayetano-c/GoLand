package main

import "fmt"

//Estructura map[tipodeclave]tipodevalor

func main() {
	visitas := make(map[string]int)
	visitas["Inicio"] = 10
	visitas["Noticias"] = 20
	visitas["Deportes"] = 25
	visitas["Hogar"] = 80
	fmt.Println(visitas)
	fmt.Println("Las Noticias:", visitas["Noticias"])
	fmt.Println("Las Noticias de hogar son:", visitas["Hogar"])
	visitas["Hogar"] = 20
	fmt.Println("Actualizado Hogar.....\nLas Noticias de hogar son:", visitas["Hogar"])
	delete(visitas, "Hogar")
	fmt.Println(visitas)

	for key, value := range visitas {
		fmt.Println(key, ":", value)
	}
}
