package main

import "fmt"

var numero1 int
var numero2 int
var secuencia int

func main() {
	//Definir los numeros de la secuencia
	numero1 := 3
	numero2 := 33

	//Realizamos la secuencia
	secuencia := numero1
	for i := 0; i < numero2; i++ {

	}
	//Mostramos el resultado
	fmt.Printf("la secuencia desde %d hasta %d da como resultado %d", numero1, numero2, secuencia)
}
