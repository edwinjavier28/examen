package main

import "fmt"

func compararNumeros(numeros []int) int {
	var secuencial int
	for i := 0; i < len(numeros); i++ {
		if numeros[i] >= secuencial {
			secuencial = numeros[i]
		}
	}
	return secuencial
}

func main() {
	numeros := []int{3, 78, 1000, 2000, 3000, 0, 4, 8}

	fmt.Println("el numero mayor es", compararNumeros(numeros))
}
