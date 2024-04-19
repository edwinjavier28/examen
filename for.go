package main

import "fmt"

var i int
var secuencia int

func main() {
	//Definir los numeros de la secuencia
	numero1 := 3
	numero2 := 10

	//Imprimimos resultados
	fmt.Printf("la secuencia entre %d y %d da como resultado", numero1, numero2)

	//Realizamos la secuencia

	for i := numero1; i <= numero2; i++ {
		//Mostramos el resultado

		fmt.Println(i)
	}
}
