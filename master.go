package main

import (
	"fmt"
	"math"
)

// Función para verificar si un número es primo
func esPrimo(numero int) bool {
	// Verificamos si el número es menor o igual a 1
	if numero <= 1 {
		return false
	}

	// Iteramos desde 2 hasta la raíz cuadrada del número
	for i := 2; i <= int(math.Sqrt(float64(numero))); i++ {
		// Si el número es divisible por algún otro número diferente de 1 y sí mismo, no es primo
		if numero%i == 0 {
			return false
		}
	}

	// Si no encontramos ningún divisor diferente de 1 y el número en sí mismo, es primo
	return true
}

func main() {
	// Número para verificar si es primo
	numero := 17

	// Verificamos si el número es primo
	if esPrimo(numero) {
		fmt.Printf("%d es un número primo\n", numero)
	} else {
		fmt.Printf("%d no es un número primo\n", numero)
	}
}
