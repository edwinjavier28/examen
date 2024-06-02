package main

import "fmt"

func main() {
	// Crear un mapa que asocia nombres con edades
	edades := map[string]int{
		"Ana":    25,
		"Luis":   30,
		"Carlos": 20,
		"Maria":  35,
	}

	// Buscar una clave en el mapa
	nombre := "Luis"
	edad, ok := edades[nombre]
	if ok {
		fmt.Printf("La edad de %s es %d.\n", nombre, edad)
	} else {
		fmt.Printf("No se encontró la edad para %s.\n", nombre)
	}

	// Intentar buscar una clave que no existe en el mapa
	nombre = "Pedro"
	edad, ok = edades[nombre]
	if ok {
		fmt.Printf("La edad de %s es %d.\n", nombre, edad)
	} else {
		fmt.Printf("No se encontró la edad para %s.\n", nombre)
	}
}
