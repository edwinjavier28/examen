package main

import "fmt"

var suma int

func main2() {
	//Se definen los numeros a sumar
	numero := 8
	numero1 := 10

	//Se realiza la suma
	suma := numero + numero1
	//Resultado de la suma
	fmt.Printf("%d mas %d da como resultado %d", numero, numero1, suma)
}
