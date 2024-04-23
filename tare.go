package main

import "fmt"

func main() {
	// Definir los elementos que quieres añadir al slice
	elementos := []int{1, 2, 3, 4, 5, 6, 7, 8}

	// Crear un slice con la longitud adecuada utilizando make
	slice := make([]int, len(elementos))

	// Copiar los elementos al slice utilizando la función copy
	copy(slice, elementos)

	// Imprimir el slice
	fmt.Println("Vector:", slice)
}
