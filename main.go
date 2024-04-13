package main

import "fmt"

func main() {
	// Número para el cual queremos imprimir la tabla de multiplicar
	for numero := 1; numero <= 9; numero++ 

	// Imprimimos la tabla de multiplicar del número
	fmt.Println("Tabla de multiplicar del", numero, ":")
	for i := 1; i <= 10; i++ {
		resultado := numero * i
		fmt.Printf("%d x %d = %d\n", numero, i, resultado)
	}

}
