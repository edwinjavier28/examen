package main

import "fmt"

// Función para calcular el factorial de un número entero
func factorial(n int) int {
	// Caso base: si n es 0 o 1, el factorial es 1
	if n == 0 || n == 1 {
		return 1
	}

	// Caso recursivo: n * factorial(n-1)
	return n * factorial(n-1)
}

func main() {
	// Número para calcular el factorial
	numero := 9

	// Calculamos el factorial del número
	resultado := factorial(numero)

	// Mostramos el resultado
	fmt.Printf("El factorial de %d es %d\n", numero, resultado)
}
