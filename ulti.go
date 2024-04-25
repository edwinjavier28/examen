package main

import "fmt"

func sumarVector1(vector []int) int {
	suma := 0
	for _, elemento := range vector {
		suma += elemento
	}
	return suma
}

func main20() {
	// Definir un vector de ejemplo
	vector := []int{19, 2, 3, 5, 10}

	// Calcular la suma de los elementos del vector
	suma := sumarVector1(vector)

	// Mostrar el resultado
	fmt.Printf("La suma de los elementos del vector es: %d", suma)
}
