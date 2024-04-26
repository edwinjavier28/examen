package main

import "fmt"

var suma int

func sumarElementos(elementos []int) int {
	suma := 0
	for _, elemento := range elementos {
		suma += elemento
	}
	return suma
}
func invertirLetras(caracteres []rune) {
	longitud := len(caracteres)
	for i := 0; i < longitud/2; i++ {
		caracteres[i], caracteres[longitud-i-1] = caracteres[longitud-i-1], caracteres[i]
	}
}

func main() {

	elementos := []int{1, 2, 3, 4, 5, 7, 8, 9, 0}
	caracteres := []rune{'c', 'a', 's', 'a'}
	slice := make([]int, len(elementos))
	slice2 := make([]rune, len(caracteres))
	copy(slice, elementos)
	copy(slice2, caracteres)
	fmt.Println(" Slice 1:\n", elementos, "\n", "Slice 2:", "\n", string(caracteres))

	suma := sumarElementos(elementos)
	fmt.Printf("La suma de los elementos del Slice 1: %d", suma)

	invertirLetras(caracteres)
	fmt.Println("\nLas letras invertidas del Slice 2:", string(caracteres))
}
