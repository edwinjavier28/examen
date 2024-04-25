package main

import "fmt"

var suma int

func sumarelementos(elementos []int) int {
	suma := 0
	for _, elemento := range elementos {
		suma += elemento
	}
	return suma
}

func main() {

	elementos := []int{1, 2, 3, 4, 5, 7, 8, 9, 0}
	caracteres := []string{"a, b, c, d, e, g, h, i, j"}
	slice := make([]int, len(elementos))
	slice2 := make([]string, len(caracteres))
	copy(slice, elementos)
	copy(slice2, caracteres)
	fmt.Println("Slice1\n", elementos, "\n", "Slice2", "\n", caracteres)

	suma := sumarelementos(elementos)
	fmt.Printf("La suma de los elementos del vector es: %d", suma)
}
