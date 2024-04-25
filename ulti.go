package main

import "fmt"

func sumarVector(vector []int) int {
	suma := 0
	for _, elemento := range vector {
		suma += elemento
	}
	return suma
}

func main() {
	// Definir un vector de ejemplo
	vector := []int{1, 2, 3, 4, 5}

	// Calcular la suma de los elementos del vector
	suma := sumarVector(vector)

	// Mostrar el resultado
	fmt.Printf("La suma de los elementos del vector es: %d\n", suma)
}
