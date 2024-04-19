package main

import "fmt"

var division int
var numero1 int
var numero2 int

func main() {
	//Se ingresan los dos numeros para la division
	numero1 := 9
	numero2 := 3

	//Se realiza la operacion
	division := numero1 / numero2

	//Se muestra el resultado
	fmt.Printf("%d dividido en %d da como resultado %d", numero1, numero2, division)
}
