package main

import "fmt"

// Función para generar los primeros n términos de la secuencia de Fibonacci
func fibonacci(n int) []int {
	// Creamos un slice para almacenar los términos de Fibonacci
	fib := make([]int, n)

	// Los dos primeros términos de Fibonacci son 0 y 1
	fib[0], fib[1] = 0, 1

	// Generamos los siguientes términos de Fibonacci
	for i := 2; i < n; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}

	return fib
}

func main() {
	// Número de términos de Fibonacci que queremos generar
	terminos := 30

	// Generamos los primeros n términos de Fibonacci
	fib := fibonacci(terminos)

	// Mostramos los términos de Fibonacci
	fmt.Printf("Los primeros %d términos de la secuencia de Fibonacci son:\n", terminos)
	fmt.Println(fib)
}
